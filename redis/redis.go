package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type (
	RedisClient interface {
		Get(ctx context.Context, key string) (interface{}, error)
		Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
		Delete(ctx context.Context, key string) error
		Redsync() *redsync.Redsync
	}

	redisInstance struct {
		client *redis.Client
	}
)

func NewRedisClient() redisInstance {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return redisInstance{client: rdb}
}

func (inst *redisInstance) Get(ctx context.Context, key string) (interface{}, error) {
	return inst.client.Get(ctx, key).Result()
}

func (inst *redisInstance) Set(ctx context.Context, key string, val interface{}, duration time.Duration) error {
	_, err := inst.client.Set(ctx, key, val, duration).Result()
	return err
}

func (inst *redisInstance) Delete(ctx context.Context, key string) error {
	_, err := inst.client.Del(ctx, key).Result()
	return err
}

func (inst *redisInstance) Redsync() *redsync.Redsync {
	pool := goredis.NewPool(inst.client)

	return redsync.New(pool)
}
