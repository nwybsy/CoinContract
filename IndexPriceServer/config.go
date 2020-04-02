package main

import (
	"CoinContract/IndexPriceServer/models/mysql"
	"CoinContract/IndexPriceServer/models/redis"
	"CoinContract/IndexPriceServer/utils"
	"errors"
	"flag"
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/util/conf"
)

type MySQLConfig struct {
	Address  string
	User     string
	Password string
	DBName   string
}

type RedisConfig struct {
	Address  string
	Password string
	DBName   int
}

func DBConfigFromFlag() map[string]string {
	resp := make(map[string]string)
	args := make(map[string]string)

	flag.Visit(func(aFlag *flag.Flag) {
		args[aFlag.Name] = aFlag.Value.String()
	})

	confPath, exist := args["config"]
	if !exist {
		panic(errors.New("config file not specified"))
	}

	c, _ := conf.NewConf(confPath)
	resp = c.GetMap("/tars/application/database")
	return resp
}

func init() {
	cfgMap := DBConfigFromFlag()

	mysqlCfg := &MySQLConfig{
		Address:  cfgMap["mysql_address"],
		User:     cfgMap["mysql_user"],
		Password: cfgMap["mysql_password"],
		DBName:   cfgMap["mysql_db_name"],
	}

	redisCfg := &RedisConfig{
		Address:  cfgMap["redis_address"],
		Password: cfgMap["redis_password"],
		DBName:   int(utils.StringToInt32(cfgMap["redis_db_name"])),
	}

	err := mysql.Init(mysqlCfg.Address, mysqlCfg.User, mysqlCfg.Password, mysqlCfg.DBName)
	fmt.Printf("mysql connect success addr:%s dbname:%s\r\n", mysqlCfg.Address, mysqlCfg.DBName)
	if err != nil {
		fmt.Printf("mysql connect error:%s\r\n", err)
		panic(err)
	}

	err = redis.Init(redisCfg.Address, redisCfg.Password, redisCfg.DBName)
	fmt.Printf("redis connect success addr:%s dbname:%d\r\n", redisCfg.Address, redisCfg.DBName)
	if err != nil {
		fmt.Printf("redis connect error:%s\r\n", err)
		panic(err)
	}
}
