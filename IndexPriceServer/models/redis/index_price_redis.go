package redis

import (
	"errors"
	"time"

	"gopkg.in/redis.v4"
)

const (
	price_t time.Duration = 6
)

func IndexPriceAll(s string) (float64, error) {
	r, err := RDB.Get(s).Float64()
	if err != nil {
		if err == redis.Nil {
			return 0.0, errors.New("redis中价格为空")
		} else {
			return 0.0, errors.New("获取redis中价格失败")
		}
	} else {
		return r, nil
	}
}

func IndexPriceAdd(s string, price float64) error {
	t := time.Second * price_t
	err := RDB.SetNX(s, price, t).Err()
	if err != nil {
		return errors.New("IndexPriceAdd redis.HSet err")
	}
	return nil
}

func IndexPriceDel(s_pre, s string) error {
	err := RDB.Del(s).Err()
	if err != nil {
		return errors.New("MinUpdateDel redis.HSet err")
	}
	return nil
}
