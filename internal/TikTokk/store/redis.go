package store

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func RedisGetWithSetNil(ctx context.Context, cli *redis.Client, redisKey string) (string, error) {
	sc := redis.NewScript(`
		local key = KEYS[1]
		local value = redis.call('GET', key)

		if not value then
			redis.call('SET', key, '')
			return ''
		else
			return value
		end
	`)
	r, err := sc.Run(ctx, cli, []string{redisKey}).Result()
	if err != nil {
		return "", err
	}
	log.Printf("[redis] get %s\n", redisKey)
	return r.(string), nil
}

func RedisSet(ctx context.Context, cli *redis.Client, redisKey string, value interface{}, time time.Duration) error {
	err := cli.Set(ctx, redisKey, value, time).Err()
	if err != nil {
		return err
	}
	log.Printf("[redis] set %s\n", redisKey)
	return nil
}

func SyncToRedis(ctx context.Context, cli *redis.Client, redisKey string, u interface{}) error {
	m, err := json.Marshal(u)
	if err != nil {
		return err
	}
	err = RedisSet(ctx, cli, redisKey, m, time.Hour)
	if err != nil {
		return err
	}
	return nil
}
