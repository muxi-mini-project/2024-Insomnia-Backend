package models

import (
	"github.com/go-redis/redis/v8"
	. "unsleepNight/config"
)

var rdb *redis.Client

func init() {
	config := LoadConfig()
	//初始化Redis客户端
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Re.Address,
		Password: config.Re.Password,
		DB:       config.Re.Database,
	})

}
