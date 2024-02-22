package models

import (
	. "Insomnia/app/common/tool"
	"Insomnia/app/request"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	TUuid string `gorm:"size:64;not null;unique"`
	Uuid  string `gorm:"size:64;not null;unique"`
	PUuid string `gorm:"size:64;not null;unique"`
	Body  string `gorm:"not null"`
	Likes uint
}

// CreatePost 方法创建一个新的回复
func CreatePost(UuID string, cp request.CreatePostReq) (post Post, err error) {
	//生成会话的TUuid
	pUuid := CreateUuid()
	post = Post{
		TUuid: cp.TUuid,
		Uuid:  UuID,
		PUuid: pUuid,
		Body:  cp.Body,
		Likes: 0,
	}
	result := Db.Create(&post)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// DestroyPost 删除指定的回复
func DestroyPost(pUuid string) error {
	return Db.Table("post").Where("p_uuid = ? ", pUuid).Delete(&Post{}).Error
}

// Posts 方法用于获取帖子的所有回复
func Posts(tUuid string) (posts []Post, err error) {
	result := Db.Where("t_uuid = ?", tUuid).Find(&posts)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// PostByPUUID 用于根据回复的PUuid查询帖子记录
func PostByPUUID(pUuid string) (post Post, err error) {
	result := Db.Where("p_uuid = ?", pUuid).First(&post)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// PostByUuId 用于根据用户的UuId 查询回复记录
func PostByUuId(UuId string) (posts []Post, err error) {
	result := Db.Where("uuid = ?", UuId).First(&posts)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func UpPostLikesData(pUuid string, exist bool) error {
	var post Post
	result := Db.Where("p_uuid = ?", pUuid).First(&post)
	if result.Error != nil {
		err := result.Error
		return err
	}
	if exist {
		post.Likes++
		return Db.Save(&post).Error
	}
	post.Likes--
	return Db.Save(&post).Error
}
