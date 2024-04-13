package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisHost     = "127.0.0.1:6379"
	redisPassword = ""
	redisDB       = 0
)

type RedisClient struct {
	*redis.Client
}

var ctx = context.Background()

func NewClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       redisDB,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return &RedisClient{rdb}
}

func (client *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return client.Client.Set(ctx, key, value, expiration).Err()
}

func (client *RedisClient) Get(key string) (string, error) {

	val, err := client.Client.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", fmt.Errorf("key does not exist")
	}
	return val, err
}
