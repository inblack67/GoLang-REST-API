package cache

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

// StartRedis ...
func StartRedis() (*redis.Client){
	redis := redis.NewClient(&redis.Options{})
	fmt.Println("Redis is here")
	return redis
}