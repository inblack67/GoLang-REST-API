package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// StartRedis ...
func StartRedis() (*redis.Client, context.Context){
	redis := redis.NewClient(&redis.Options{})
	fmt.Println("Redis is here")

	ctx := context.Background()

	redis.FlushAll(ctx)

	redis.Set(ctx, "hello", "worlds", 0)

	return redis, ctx
}

// SET ...
func SET(key string, value string) error{
	redisClient, ctx := StartRedis()
	err := redisClient.Set(ctx, key, value, 0).Err()
	return err
}

// GET ...
func GET(key string) (string, error){
	redisClient, ctx := StartRedis()
	val , err := redisClient.Get(ctx, key).Result()
	return val, err
}