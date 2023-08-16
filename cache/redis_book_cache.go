package cache

import (
	"context"
	"encoding/json"
	"golang_rest_api/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	redisClient *redis.Client
	expire      time.Duration
	ctx         context.Context
}

func NewRedisCache(host string, db int, exp time.Duration, password string, ctx context.Context) BookCache {
	return &RedisCache{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password,
			DB:       db,
		}),
		expire: exp,
		ctx:    ctx,
	}
}

func (cache *RedisCache) Set(key *string, value *models.Book) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	cache.redisClient.Set(cache.ctx, *key, json, cache.expire*time.Second)
	return nil
}

func (cache *RedisCache) Get(key *string) (*models.Book, error) {
	val, err := cache.redisClient.Get(cache.ctx, *key).Result()
	if err != nil {
		return nil, err
	}

	book := models.Book{}
	err = json.Unmarshal([]byte(val), &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
