package lib

import (
	"os"
	"strconv"
)

type Env struct {
	ServerPort    string
	LogOutput     string
	DBUsername    string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	CacheHost     string
	CachePort     string
	CachePassword string
	ApiKey        string
	ApiRate       int
}

func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

func (env *Env) LoadEnv() {
	env.ServerPort = os.Getenv("APP_PORT")
	env.DBUsername = os.Getenv("DB_USERNAME")
	env.DBPassword = os.Getenv("DB_PASSWORD")
	env.DBHost = os.Getenv("DB_HOST")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBName = os.Getenv("DB_NAME")
	env.CacheHost = os.Getenv("REDIS_HOST")
	env.CachePort = os.Getenv("REDIS_PORT")
	env.CachePassword = os.Getenv("REDIS_PASSWORD")
	env.ApiKey = os.Getenv("API_KEY")
	env.ApiRate, _ = strconv.Atoi(os.Getenv("API_RATE"))
}
