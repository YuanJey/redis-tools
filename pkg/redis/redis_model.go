package redis

import (
	"context"
	"github.com/YuanJey/redis-tools/pkg/config"
	go_redis "github.com/go-redis/redis/v8"
	"time"
)

const (
	uidPidToken     = "UID_PID_TOKEN_STATUS:"
	accountTempCode = "ACCOUNT_TEMP_CODE"
)

type Redis struct {
	RDB go_redis.UniversalClient
}

var RedisDB *Redis

func (d *Redis) SetAccountCode(account string, code, ttl int) (err error) {
	key := accountTempCode + account
	return d.RDB.Set(context.Background(), key, code, time.Duration(ttl)*time.Second).Err()
}
func (d *Redis) GetAccountCode(account string) (string, error) {
	key := accountTempCode + account
	return d.RDB.Get(context.Background(), key).Result()
}
func (d *Redis) JudgeAccountEXISTS(account string) (bool, error) {
	key := accountTempCode + account
	n, err := d.RDB.Exists(context.Background(), key).Result()
	if n > 0 {
		return true, err
	} else {
		return false, err
	}
}
func init() {
	RedisDB = &Redis{}
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	RedisDB.RDB = go_redis.NewClient(&go_redis.Options{
		Addr:     config.ServerConfig.Redis.DBAddress,
		Username: config.ServerConfig.Redis.DBUserName,
		Password: config.ServerConfig.Redis.DBPassWord, // no password set
		DB:       0,                                    // use default DB
		PoolSize: 100,                                  // 连接池大小
	})
	//_, err := RedisDB.RDB.Ping(ctx).Result()
	//if err != nil {
	//	panic(err.Error() + " redis " + config.Config.Redis.DBAddress[0] + config.Config.Redis.DBUserName + config.Config.Redis.DBPassWord)
	//}
}
