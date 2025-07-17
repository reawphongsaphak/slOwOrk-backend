package config

import (
	"os"
	"sync"
)

var config Config
var loadOnce sync.Once

type Config struct {
	Port       	string
	DBUrl      	string
	DB_NAME		string
	JWTSecret 	string
}

func LoadConfig() Config {
	loadOnce.Do(func() {
		config = Config{
			Port: os.Getenv("PORT"),
			DBUrl: os.Getenv("MONGO_URI"),
			DB_NAME: os.Getenv("DB_NAME"),
			JWTSecret: os.Getenv("JWT_SECRET"),
			}
		},
	)
	return config
}