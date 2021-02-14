package cache

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

// StartRedis ...
func StartRedis() (*redis.Client){
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
	})
	fmt.Println("Redis is here")
	return redis
}