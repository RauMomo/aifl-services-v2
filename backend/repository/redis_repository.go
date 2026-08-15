package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"aifl-backend/common/config"
	"aifl-backend/common/logger"

	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	IsActive() bool
	CacheSet(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	CacheGet(ctx context.Context, key string, target interface{}) bool
	CacheDelete(ctx context.Context, keys ...string)
	RateLimit(ctx context.Context, key string, maxHits int64, window time.Duration) (bool, error)
}

type redisRepository struct {
	client *redis.Client
	active bool
}

func NewRedisRepository(cfg *config.Config) RedisRepository {
	var opts *redis.Options

	redisURL := ""
	host := "localhost"
	port := "6379"
	password := ""

	if cfg != nil {
		redisURL = cfg.Redis.URL
		if cfg.Redis.Host != "" {
			host = cfg.Redis.Host
		}
		if cfg.Redis.Port != "" {
			port = cfg.Redis.Port
		}
		password = cfg.Redis.Password
	}

	if redisURL != "" {
		parsedOpts, err := redis.ParseURL(redisURL)
		if err != nil {
			logger.Log.Warn("Invalid REDIS_URL format, falling back to host:port", "error", err.Error())
			opts = &redis.Options{Addr: fmt.Sprintf("%s:%s", host, port), Password: password}
		} else {
			opts = parsedOpts
		}
	} else {
		opts = &redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       0,
		}
	}

	client := redis.NewClient(opts)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		logger.Log.Warn("🔴 Redis server not reachable. Running with Redis Cache Fallback (In-Memory Direct Mode)", "addr", opts.Addr)
		return &redisRepository{client: nil, active: false}
	}

	logger.Log.Info("🟢 Connected to Redis Server successfully", "addr", opts.Addr)
	return &redisRepository{client: client, active: true}
}

func (r *redisRepository) IsActive() bool {
	return r.active && r.client != nil
}

func (r *redisRepository) CacheSet(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	if !r.IsActive() {
		return nil
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, ttl).Err()
}

func (r *redisRepository) CacheGet(ctx context.Context, key string, target interface{}) bool {
	if !r.IsActive() {
		return false
	}
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	return json.Unmarshal([]byte(val), target) == nil
}

func (r *redisRepository) CacheDelete(ctx context.Context, keys ...string) {
	if !r.IsActive() {
		return
	}
	r.client.Del(ctx, keys...)
}

func (r *redisRepository) RateLimit(ctx context.Context, key string, maxHits int64, window time.Duration) (bool, error) {
	if !r.IsActive() {
		return true, nil // Allow request if Redis is offline
	}

	pipe := r.client.TxPipeline()
	incr := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, window)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return true, err
	}

	return incr.Val() <= maxHits, nil
}
