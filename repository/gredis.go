package redis

import (
	cfg "auto-installation/config"
	"fmt"
	"github.com/go-redis/redis"
)

var Client *redis.Client

func InitRedisClient() {
	Client = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%d", cfg.Cfg.TaskRedisHost, cfg.Cfg.TaskRedisPort),
		Password:           cfg.Cfg.TaskRedisPassword,
		DB:                 cfg.Cfg.TaskRedisDB,
		PoolSize:           cfg.Cfg.TaskRedisPoolSize,
	})
}
