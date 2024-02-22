package models

import (
	"gorm.io/gorm"
)

type Likes struct {
	gorm.Model
	Uid   string `gorm:"size:64;not null"`
	Uuid  string `gorm:"size:64;not null"`
	Exist bool   `gorm:"not null;unique"`
}

func ChangeLike(uid, uuid string) (exist bool, err error) {
	// 在数据库中查找是否存在对应的记录
	like := Likes{
		Uid:   uid,
		Uuid:  uuid,
		Exist: false,
	}

	result := Db.Where("uid = ? AND uuid = ?", uid, uuid).FirstOrCreate(&like)
	if result.Error != nil {
		return false, result.Error
	}

	// 更改点赞状态
	like.Exist = !like.Exist

	// 保存更改
	result = Db.Save(&like)
	if result.Error != nil {
		return false, result.Error
	}

	return like.Exist, nil
}

func CheckLike(uuid string, uid string) (exist bool, err error) {
	var l Likes
	err = Db.Table("likes").Where("uuid = ? AND uid = ?", uuid, uid).First(&l).Error
	if err != nil {
		exist = false
		return
	}
	exist = l.Exist
	err = nil
	return
}
