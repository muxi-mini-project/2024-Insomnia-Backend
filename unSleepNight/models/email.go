package models

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type CheckEmail struct {
	Email            string `gorm:"size:255;not null;unique"`
	VerificationCode string `gorm:"size:255;not null"`
}

// CreateRedis 创建一个Redis
func (email CheckEmail) CreateRedis() (err error) {
	email.VerificationCode = Encrypt(email.VerificationCode)
	ctx := context.Background()
	key := "verification_code:" + email.Email
	expiration := 5 * time.Minute
	err = rdb.Set(ctx, key, email.VerificationCode, expiration).Err()
	if err != nil {
		return
	}
	return
}

// CheckVerificationCode 方法检查验证码是否正确
func (email CheckEmail) CheckVerificationCode() (err error) {
	email.VerificationCode = Encrypt(email.VerificationCode)
	// 在 Redis 中检查验证码
	ctx := context.Background()
	key := "verification_code:" + email.Email
	storedCode, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	if storedCode != email.VerificationCode {
		return errors.New("无效的验证码")
	}
	return
}

// CheckIfEmailExists 检查是否存在email,存在就不让注册
func CheckIfEmailExists(email string) (exists bool, err error) {
	//根据email查询用户记录
	result := Db.Where("email = ?", email).First(&User{})
	if result.Error == nil {
		// 记录已找到，说明邮箱已经存在
		return true, nil
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 记录未找到，说明邮箱未注册
		return false, nil
	}
	// 其他错误情况，返回错误
	err = result.Error
	return
}
