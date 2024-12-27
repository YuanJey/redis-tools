package redis

import (
	"github.com/YuanJey/redis-tools/pkg/config"
	go_redis "github.com/go-redis/redis/v8"
)

type Redis struct {
	RDB go_redis.UniversalClient
}

var RedisDB *Redis

func init() {
	RedisDB = &Redis{}
	RedisDB.RDB = go_redis.NewClient(&go_redis.Options{
		Addr:     config.ServerConfig.Redis.DBAddress,
		Username: config.ServerConfig.Redis.DBUserName,
		Password: config.ServerConfig.Redis.DBPassWord, // no password set
		DB:       0,                                    // use default DB
		PoolSize: 100,                                  // 连接池大小
	})
}
