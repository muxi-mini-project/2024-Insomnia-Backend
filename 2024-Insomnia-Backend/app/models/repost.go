package models

import (
	. "Insomnia/app/common/tool"
	"Insomnia/app/request"
	"gorm.io/gorm"
)

type RePost struct {
	gorm.Model
	TUuid string `gorm:"size:64;not null;unique"`
	Uuid  string `gorm:"size:64;not null;unique"`
	PUuid string `gorm:"size:64;not null;unique"`
	RUuid string `gorm:"size:64;not null;unique"`
	Body  string `gorm:"not null"`
	Likes uint
}

// CreateRePost 方法创建一个新的re回复
func CreateRePost(UuID string, crp request.CreateRePostReq) (repost RePost, err error) {
	//生成会话的RUuid
	rUuid := CreateUuid()
	repost = RePost{
		TUuid: crp.TUuid,
		Uuid:  UuID,
		PUuid: crp.PUuid,
		RUuid: rUuid,
		Body:  crp.Body,
		Likes: 0,
	}
	result := Db.Create(&repost)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// DestroyRePost 删除指定的re回复
func DestroyRePost(rUuid string) error {
	return Db.Table("repost").Where("r_uuid = ? ", rUuid).Delete(&RePost{}).Error
}

// RePostByRUUID 用于根据回复的RUuid查询帖子记录
func RePostByRUUID(rUuid string) (repost RePost, err error) {
	result := Db.Where("r_uuid = ?", rUuid).First(&repost)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// RePosts 方法用于获取回复的所有re回复
func RePosts(rUuid string) (reposts []RePost, err error) {
	result := Db.Where("r_uuid = ?", rUuid).Find(&reposts)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// RepostByUuId 用于根据用户的UuId 查询回复记录
func RepostByUuId(UuId string) (reposts []RePost, err error) {
	result := Db.Where("uuid = ?", UuId).First(&reposts)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func UpRePostLikesData(pUuid string, exist bool) error {
	var rePost RePost
	result := Db.Where("r_uuid = ?", pUuid).First(&rePost)
	if result.Error != nil {
		err := result.Error
		return err
	}
	if exist {
		rePost.Likes++
		return Db.Save(&rePost).Error
	}
	rePost.Likes--
	return Db.Save(&rePost).Error
}
