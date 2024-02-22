package main

import (
	"Insomnia/app/core/middlewares"
	"Insomnia/app/routers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

// 个人总数据
type taskSumData struct {
	gorm.Model
	Uuid  string `gorm:"size:64;not null;unique"`
	Email string `gorm:"size:255;not null;unique"`
	Sum   uint
	Tasks string
}

// @title 不眠夜
// @version 1.0
// @description 一个匿名熬夜论坛
func main() {
	//启动gin的engine
	engine := gin.Default()
	routers.Load(engine)
	//加载中间件
	middlewares.Load(engine)

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
