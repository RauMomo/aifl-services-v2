package db

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"aifl-backend/config"
	"aifl-backend/logger"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	redisActive bool
)

// InitRedis initializes the Redis client or sets graceful fallback mode
func InitRedis() {
	var opts *redis.Options

	redisURL := ""
	host := "localhost"
	port := "6379"
	password := ""

	if config.AppConfig != nil {
		redisURL = config.AppConfig.Redis.URL
		if config.AppConfig.Redis.Host != "" {
			host = config.AppConfig.Redis.Host
		}
		if config.AppConfig.Redis.Port != "" {
			port = config.AppConfig.Redis.Port
		}
		password = config.AppConfig.Redis.Password
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
		redisActive = false
		return
	}

	RedisClient = client
	redisActive = true
	logger.Log.Info("🟢 Connected to Redis Server successfully", "addr", opts.Addr)
}

// IsRedisActive returns true if Redis is online
func IsRedisActive() bool {
	return redisActive && RedisClient != nil
}

// CacheSet serializes and stores an object in Redis with expiration
func CacheSet(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	if !IsRedisActive() {
		return nil
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return RedisClient.Set(ctx, key, data, ttl).Err()
}

// CacheGet retrieves and deserializes a Redis cached item
func CacheGet(ctx context.Context, key string, target interface{}) bool {
	if !IsRedisActive() {
		return false
	}
	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	return json.Unmarshal([]byte(val), target) == nil
}

// CacheDelete removes keys matching a given pattern or key
func CacheDelete(ctx context.Context, keys ...string) {
	if !IsRedisActive() {
		return
	}
	RedisClient.Del(ctx, keys...)
}

// RateLimit checks if an IP/key exceeds allowed hits within a window duration
func RateLimit(ctx context.Context, key string, maxHits int64, window time.Duration) (bool, error) {
	if !IsRedisActive() {
		return true, nil // Allow request if Redis is offline
	}

	pipe := RedisClient.TxPipeline()
	incr := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, window)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return true, err
	}

	return incr.Val() <= maxHits, nil
}
