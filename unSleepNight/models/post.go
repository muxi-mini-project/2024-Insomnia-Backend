package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Uuid      string `gorm:"size:64;not null;unique"`
	Body      string
	UserId    uint      `gorm:"not null;unique"`
	ThreadId  uint      `gorm:"not null;unique"`
	CreatedAt time.Time `gorm:"type:datetime;default:null"`
}

// CreatedAtDate 方法获取创建时间
func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("2006-01-02 15:04")
}

// User 获取创建者
func (post *Post) User() (user User) {
	result := Db.First(&user, post.UserId)
	if result.Error != nil {
		return
	}
	return
}
