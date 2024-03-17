package models

import (
	. "Insomnia/app/core/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os/exec"
)

var rdb *redis.Client

func init() {
	// 检查 Redis 服务器是否已经运行
	redisRunning, err := isRedisRunning()
	if err != nil {
		fmt.Println("Failed to check Redis status:", err)
		return
	}
	// 如果 Redis 服务器未运行，则尝试启动 Redis 服务器
	if !redisRunning {
		err := startRedisServer()
		if err != nil {
			fmt.Println("Failed to start Redis server:", err)
			return
		}
		fmt.Println("Redis server started.")
		config := LoadConfig()
		//初始化Redis客户端
		rdb = redis.NewClient(&redis.Options{
			Addr:     config.Re.Address,
			Password: config.Re.Password,
			DB:       config.Re.Database,
		})
	}
}

// 检查 Redis 服务器是否已经运行
func isRedisRunning() (bool, error) {
	ctx := context.Background()
	config := LoadConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     config.Re.Address,
		Password: config.Re.Password,
		DB:       config.Re.Database,
	})
	defer client.Close()
	_, err := client.Ping(ctx).Result()
	if err == nil {
		return true, nil
	}
	if err.Error() == "redis: can't connect to the server" {
		return false, err
	}
	return false, nil
}

// 启动 Redis 服务器
func startRedisServer() error {
	cmd := exec.Command("redis-server")
	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil
}
