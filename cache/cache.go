package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

// StartRedis ...
func StartRedis() (*redis.Client){
	redis := redis.NewClient(&redis.Options{})
	fmt.Println("Redis is here")

	ctx := context.Background()

	redis.FlushAll(ctx)

	err := redis.Set(ctx, "hello", "worlds", 0).Err()
	if err != nil{
		log.Fatal(err)
	}

	value, err2 := redis.Get(ctx, "hello").Result()

	if err2 != nil{
		log.Fatal(err)
	}

	fmt.Println("===========", value)

	defer redis.Close()

	return redis
}