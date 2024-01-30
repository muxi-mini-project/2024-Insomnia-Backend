package models

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Db 连接数据库的指针
var Db *gorm.DB

// 数据库连接启动!
func init() {
	var err error
	user := "root"
	password := "Sjn3265926531"
	dbname := "test"
	//打开数据库
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbname)
	//检测数据库连接是否正常
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// 创建用户数据库表
	err = Db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
	return
}

// 创建uuid
func createUuid() string {
	u := make([]byte, 16)
	_, err := rand.Read(u)
	if err != nil {
		log.Fatal("Cannot generate Uuid", err)
	}

	//看不懂的随机生成部分(看懂了是位运算)
	//0x40 是RFC 4122 中保留的变体
	u[8] = (u[8] | 0x40) & 0x7F
	//将time_hi_and_version字段的四个最高有效位(第12到15位),设置为4位版本号
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

// Encrypt 使用SHA-1对明文密码进行哈希加密
func Encrypt(plaintext string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
}
