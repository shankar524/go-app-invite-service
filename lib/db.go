package lib

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(env Env) Database {
	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbName := env.DBName
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Println("Url: ", url)
		log.Panic(err)
	}
	log.Println("Database connection established")
	return Database{
		DB: db,
	}
}
