package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"strconv"
	"time"
)

var redisdb *redis.Client

func GetRedis() *redis.Client {
	return redisdb
}

func init() {
	redisInit()
	log.Println("[INIT REDIS CONNECTS] success")
}

func redisInit() {

	host := os.Getenv("REDIS_HOSTNAME")
	port := os.Getenv("REDIS_PORT")
	dbStr := os.Getenv("REDIS_SELECT")
	passwd := os.Getenv("REDIS_PASSWORD")

	var db int
	addr := host + ":" + port
	if dbStr == "" {
		db = 0
	} else {
		db, _ = strconv.Atoi(dbStr)
	}

	redisdb = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     passwd,
		DB:           db,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     50,
		PoolTimeout:  30 * time.Second,
	})

	_, err := redisdb.Ping(context.Background()).Result()
	if err != nil {
		log.Panicln(err.Error())
	}

}
