package models

import (
	. "Insomnia/app/common/tool"
	"Insomnia/app/request"
	"encoding/json"
	"gorm.io/gorm"
	"log"
)

type Thread struct {
	gorm.Model
	TUuid      string `gorm:"size:64;not null;unique"`
	Topic      string `gorm:"not null"`
	Title      string `gorm:"size:64;not null;"`
	Uuid       string `gorm:"not null"`
	Likes      uint
	Body       string `gorm:"not null"`
	PostNumber uint
	Images     string `gorm:"type:json"`
}

// CreateThread 方法创建一个新的帖子
func CreateThread(UuiD string, ct request.CreateThreadReq) (thread Thread, err error) {
	//生成会话的TUuid
	tUuid := CreateUuid()
	jsonData, err := json.Marshal(ct.Images)
	if err != nil {
		log.Fatal("序列化数据失败:", err)
	}
	thread = Thread{
		Model:      gorm.Model{},
		TUuid:      tUuid,
		Topic:      ct.Topic,
		Uuid:       UuiD,
		Body:       ct.Body,
		Likes:      0,
		PostNumber: 0,
		Title:      ct.Title,
		Images:     string(jsonData),
	}
	result := Db.Create(&thread)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// DestroyThread 删除指定的帖子
func DestroyThread(tUuid string) (err error) {
	err = Db.Table("threads").Where("t_uuid = ? ", tUuid).Delete(&Thread{}).Error
	if err != nil {
		return
	}
	err = Db.Table("posts").Where("t_uuid = ? ", tUuid).Delete(&Thread{}).Error
	if err != nil {
		return
	}
	err = Db.Table("re_posts").Where("t_uuid = ? ", tUuid).Delete(&Thread{}).Error
	return
}

// NumReplies 方法用于获取帖子的回复数量
func (thread *Thread) NumReplies() (count int64) {
	result := Db.Model(&Post{}).Where("t_uuid = ?", thread.ID).Count(&count)
	if result.Error != nil {
		return
	}
	return
}

// Posts 方法用于获取该帖子的所有回复
func (thread *Thread) Posts() (posts []Post, err error) {
	result := Db.Where("t_uuid = ?", thread.TUuid).Find(&posts)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// Threads 方法用于获取数据库中的所有帖子记录
func Threads(topic string) (threads []Thread, err error) {
	result := Db.Where("topic = ?", topic).Order("created_at desc").Find(&threads)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// ThreadByTUUID 用于根据帖子的TUuid查询帖子记录
func ThreadByTUUID(tUuid string) (thread Thread, err error) {
	result := Db.Where("t_uuid = ?", tUuid).First(&thread)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// ThreadByUuId 用于根据用户的UuId 查询帖子记录
func ThreadByUuId(UuId string) (threads []Thread, err error) {
	result := Db.Where("uuid = ?", UuId).First(&threads)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func UpThreadLikesData(tUuid string, exist bool) error {
	var thread Thread
	result := Db.Where("t_uuid = ?", tUuid).First(&thread)
	if result.Error != nil {
		err := result.Error
		return err
	}
	if exist {
		thread.Likes++
		return Db.Save(&thread).Error
	}
	thread.Likes--
	return Db.Save(&thread).Error
}
