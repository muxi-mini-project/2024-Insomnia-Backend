package models

import (
	"gorm.io/gorm"
	"time"
)

type Thread struct {
	gorm.Model
	Uuid      string `gorm:"size:64;not null;unique"`
	Topic     string `gorm:"not null"`
	UserId    uint   `gorm:"not null;unique"`
	likes     uint
	CreatedAt time.Time `gorm:"type:datetime;default:null"`
}

// CreatedAtDate 方法获取创建的时间
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("2006-01-02 15:04")
}

// NumReplies 方法用于获取帖子的回复数量
func (thread *Thread) NumReplies() (count int64) {
	result := Db.Model(&Post{}).Where("thread_id = ?", thread.ID).Count(&count)
	if result.Error != nil {
		return
	}
	return
}

// Posts 方法用于获取帖子的所有回复
func (thread *Thread) Posts() (posts []Post, err error) {
	result := Db.Where("thread_id = ?", thread.ID).Find(&posts)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// Threads 方法用于获取数据库中的所有帖子记录
func Threads() (threads []Thread, err error) {
	result := Db.Order("created_at desc,likes desc").Find(&threads)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// ThreadByUUID 方法用于根据 UUID 查询帖子记录
func ThreadByUUID(uuid string) (thread Thread, err error) {
	result := Db.Where("uuid = ?", uuid).First(&thread)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// User 方法用于获取创建该帖子的用户信息
func (thread *Thread) User() (user User) {
	result := Db.First(&user, thread.UserId)
	if result.Error != nil {
		return
	}
	return
}
