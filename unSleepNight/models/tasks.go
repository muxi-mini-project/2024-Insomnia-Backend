package models

import (
	"gorm.io/gorm"
)

// Task 个人总数据
type Task struct {
	gorm.Model
	Uuid string `gorm:"size:64;not null;unique"`
	Sum  uint
	Day  int `gorm:"not null"`
}

// Update 方法更新数据+1
func (task Task) Update() (err error) {
	task.Sum += 1
	result := Db.Model(task).Update("sum", task.Sum)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
