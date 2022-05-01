package lib

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type Cache struct {
	Connection *redis.Pool
}

type ICache interface {
	Save(string) error
	Exists(string) (bool, error)
	Delete(string) (bool, error)
}

func NewCache(env Env) ICache {
	host := env.CacheHost
	port := env.CachePort
	password := env.CachePassword

	RedisConn := &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
			if err != nil {
				return nil, err
			}

			if password != "" {
				if _, err = c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},
	}
	log.Print("Cache connection established")
	return &Cache{RedisConn}
}

// Set a key/value
func (c Cache) Save(key string) error {
	log.Println("cache:: save")

	conn := c.Connection.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, true)
	if err != nil {
		return err
	}

	return nil
}

// Get gets a key
func (c Cache) Exists(key string) (bool, error) {
	log.Println("cache:: exists")

	conn := c.Connection.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("EXISTS", key))
}

// Delete delete a key
func (c Cache) Delete(key string) (bool, error) {
	log.Println("cache:: Delete")

	conn := c.Connection.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}
