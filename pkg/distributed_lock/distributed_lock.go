package distributed_lock

import (
	"context"
	"fmt"
	"github.com/YuanJey/redis-tools/pkg/redis"
	"time"
)

var ctx = context.Background()

// DistributedLock 表示一个分布式锁
type DistributedLock struct {
	client *redis.Redis
	key    string
	value  string
	ttl    time.Duration
}

// NewDistributedLock 返回一个新的分布式锁实例
func NewDistributedLock(client *redis.Redis, key string, value string, ttl time.Duration) *DistributedLock {
	return &DistributedLock{
		client: client,
		key:    key,
		value:  value,
		ttl:    ttl,
	}
}

// Acquire 尝试获取锁
func (l *DistributedLock) Acquire() bool {
	result, err := l.client.RDB.SetNX(ctx, l.key, l.value, l.ttl).Result()
	if err != nil {
		fmt.Println("获取锁时发生错误:", err)
		return false
	}
	return result
}

// Release 释放锁
func (l *DistributedLock) Release() bool {
	// 使用 Lua 脚本确保原子性
	script := `  
    if redis.call("get", KEYS[1]) == ARGV[1] then  
        return redis.call("del", KEYS[1])  
    else  
        return 0  
    end`
	result, err := l.client.RDB.Eval(ctx, script, []string{l.key}, l.value).Result()
	if err != nil {
		fmt.Println("释放锁时发生错误:", err)
		return false
	}
	return result.(int64) == 1
}
