package main

import (
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

// getRedisOptions will get the redis options from env variables
func getRedisOptions() *redis.Options {
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		log.Fatalln("please set REDIS_URL environmental variable")
	}
	dbID, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	username := os.Getenv("REDIS_USERNAME")
	password := os.Getenv("REDIS_PASSWORD")
	return &redis.Options{
		Addr:     redisUrl,
		DB:       dbID,
		Username: username,
		Password: password,
	}
}

func getDatabaseDSN() string {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatalln("please set DATABASE_DSN env variable")
	}
	return dsn
}

func getListenAddress() string {
	address := os.Getenv("LISTEN_ADDRESS")
	if address == "" {
		log.Fatalln("please set LISTEN_ADDRESS env variable")
	}
	return address
}
