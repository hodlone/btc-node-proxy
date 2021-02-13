package cache

import (
	"context"
	"log"

	"github.com/NodeHodl/btc-node-proxy/config"

	"github.com/go-redis/redis/v8"
)

// Redis ...
type Redis struct {
	*redis.Client
}

var ctx = context.Background()

// NewRedisClient ...
func NewRedisClient() *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Pass,
	})

	return &Redis{
		client,
	}
}

// Get ...
func (r *Redis) Get(key string) (string, error) {
	var value string
	var err error

	if value, err = r.Client.Get(ctx, key).Result(); err != nil {
		log.Printf("[REDIS GET() ERROR]\n%v", err)
		return "", err
	}

	return value, nil
}

// Set ...
func (r *Redis) Set(key, value string) error {
	var err error

	if err = r.Client.Set(ctx, key, value, 0).Err(); err != nil {
		log.Printf("[REDIS SET() ERROR]\n%v", err)
		return err
	}

	return nil
}
