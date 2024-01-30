package models

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	Uuid      string    `gorm:"size:64;not null;unique"`
	Name      string    `gorm:"size:255;not null;unique"`
	Email     string    `gorm:"size:255;not null;unique"`
	UserId    uint      `gorm:"not null;unique"`
	CreatedAt time.Time `gorm:"type:datetime;default:null"`
}

// Check 方法检查数据库中是否有该session
func (session *Session) Check() (valid bool, err error) {
	result := Db.Where("uuid = ?", session.Uuid).First(&session)
	if result.Error != nil {
		err = result.Error
		return
	}
	valid = result.RowsAffected > 0
	return
}

// DeleteByUuid 方法通过uuid删除session
func (session *Session) DeleteByUuid() (err error) {
	result := Db.Where("uuid = ?", session.Uuid).Delete(&Session{})
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// SessionDeleteAll 方法删除所有的session
func SessionDeleteAll() (err error) {
	result := Db.Delete(&Session{})
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
