package common

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/Laisky/zap"
	"github.com/go-redis/redis/v8"

	"github.com/songquanpeng/one-api/common/config"
	"github.com/songquanpeng/one-api/common/logger"
)

var RDB redis.Cmdable
var RedisEnabled = true

// InitRedisClient This function is called after init()
func InitRedisClient() (err error) {
	if os.Getenv("REDIS_CONN_STRING") == "" {
		RedisEnabled = false
		logger.Logger.Info("REDIS_CONN_STRING not set, Redis is not enabled")
		return nil
	}
	if config.SyncFrequency == 0 {
		RedisEnabled = false
		logger.Logger.Info("SYNC_FREQUENCY not set, Redis is disabled")
		return nil
	}
	redisConnString := os.Getenv("REDIS_CONN_STRING")
	if os.Getenv("REDIS_MASTER_NAME") == "" {
		logger.Logger.Info("Redis is enabled")
		opt, err := redis.ParseURL(redisConnString)
		if err != nil {
			logger.Logger.Fatal("failed to parse Redis connection string", zap.Error(err))
		}
		RDB = redis.NewClient(opt)
	} else {
		// cluster mode
		logger.Logger.Info("Redis cluster mode enabled")
		RDB = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:      strings.Split(redisConnString, ","),
			Password:   os.Getenv("REDIS_PASSWORD"),
			MasterName: os.Getenv("REDIS_MASTER_NAME"),
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		logger.Logger.Fatal("Redis ping test failed", zap.Error(err))
	}
	return err
}

func ParseRedisOption() *redis.Options {
	opt, err := redis.ParseURL(os.Getenv("REDIS_CONN_STRING"))
	if err != nil {
		logger.Logger.Fatal("failed to parse Redis connection string", zap.Error(err))
	}
	return opt
}

func RedisSet(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	return RDB.Set(ctx, key, value, expiration).Err()
}

func RedisGet(key string) (string, error) {
	ctx := context.Background()
	return RDB.Get(ctx, key).Result()
}

func RedisDel(key string) error {
	ctx := context.Background()
	return RDB.Del(ctx, key).Err()
}

func RedisDecrease(key string, value int64) error {
	ctx := context.Background()
	return RDB.DecrBy(ctx, key, value).Err()
}
