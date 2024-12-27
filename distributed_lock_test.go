package redis_tools

import (
	"fmt"
	"github.com/YuanJey/redis-tools/pkg/distributed_lock"
	"github.com/YuanJey/redis-tools/pkg/redis"
	"testing"
	"time"
)

func TestNewDistributedLock(t *testing.T) {
	lock := distributed_lock.NewDistributedLock(redis.RedisDB, "my_lock", "unique_value", 5*time.Second)

	// 尝试获取锁
	if lock.Acquire() {
		fmt.Println("锁已获取!")

		// 执行一些工作...
		time.Sleep(2 * time.Second)

		// 释放锁
		if lock.Release() {
			fmt.Println("锁已释放!")
		} else {
			fmt.Println("释放锁失败!")
		}
	} else {
		fmt.Println("获取锁失败!")
	}
}
