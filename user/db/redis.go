package db

import "github.com/go-redis/redis/v8"

func CreateRedisConnection(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   dbNo,
	})

	return rdb
}
