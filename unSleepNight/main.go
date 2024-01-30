package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	. "unsleepNight/models"
	. "unsleepNight/routes"
)

// 个人总数据
type taskSumData struct {
	gorm.Model
	Uuid  string `gorm:"size:64;not null;unique"`
	Email string `gorm:"size:255;not null;unique"`
	Sum   uint
	Tasks string
}

func main() {
	user := "root"
	password := "Sjn3265926531"
	dbname := "test"
	//打开数据库
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbname)

	//定义err
	var err error

	//检测数据库连接是否正常
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}

	// 创建任务数据库表
	err = Db.AutoMigrate(&Post{})
	if err != nil {
		panic("无法创建Post数据库表")
	}

	// 创建用户数据库表
	err = Db.AutoMigrate(&User{})
	if err != nil {
		panic("无法创建User数据库表")
	}

	// 创建用户数据库表
	err = Db.AutoMigrate(&Session{})
	if err != nil {
		panic("无法创建Session数据库表")
	}

	// 创建用户数据库表
	err = Db.AutoMigrate(&Thread{})
	if err != nil {
		panic("无法创建Thread数据库表")
	}

	//启动服务192.168.226.121
	err = startWebServer(":8080")
	if err != nil {
		return
	}
}

// 检查是否已经登陆
func checkIfLogin(context *gin.Context) (string, error) {

	//获取 Cookie 中的用户信息,检查是否已经登陆
	Email, err := context.Cookie("email")
	if err != nil || Email == "" {
		// Cookie 中无用户信息，说明用户未登录
		context.String(http.StatusOK, "您还未登陆,请先登陆")
		return Email, err
	}
	return Email, err
}

// 通过制定端口启动web服务器通过改变port来实现
func startWebServer(port string) error {
	r := NewRouter()
	//处理静态资源
	r.Static("/static", "./public")
	//端口启动!
	err := r.Run(port)
	if err != nil {
		fmt.Println("服务器启动失败")
		return err
	}
	return err
}
