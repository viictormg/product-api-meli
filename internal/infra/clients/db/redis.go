package db

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/viictormg/product-api-meli/config"
)

var ctx = context.Background()

func NewRedisConnection(config *config.Config) *redis.Client {
	configRedis := config.GetRedisConfig()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configRedis.Host, configRedis.Port),
		Password: "", //
		DB:       0,  //
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}

	fmt.Println("Connected to Redis:", pong)

	return client
}
