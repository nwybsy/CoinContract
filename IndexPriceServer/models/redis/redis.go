package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

var RDB *redis.Client

func Init(address, password string, DBName int) error {
	RDB = redis.NewClient(&redis.Options{
		Addr:         address,
		Password:     password,
		DB:           DBName,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	_, err := RDB.Ping().Result()
	if err != nil {
		fmt.Println("rdb.Ping err:", err)
		return err
	}
	return nil
}

func AcquireLock(lockKey string) {
	ttl := time.Second * 30
	ok, err := RDB.SetNX(lockKey, 1, ttl).Result()
	for err != nil || !ok {
		ok, err = RDB.SetNX(lockKey, 1, ttl).Result()
	}
}

func ReleaseLock(lockKey string) error {
	return RDB.Del(lockKey).Err()
}
