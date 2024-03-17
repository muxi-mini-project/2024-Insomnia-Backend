package models

import (
	"gorm.io/gorm"
	"time"
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

// GetAllTaskByUuid 方法获取所有的task数据
func GetAllTaskByUuid(uuid string) (task []Task, err error) {
	//根据email查询用户记录
	result := Db.Where("uuid = ?", uuid).First(&task)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// GetTaskByUuid 方法获取当天的数据，不存在会自动创建
func GetTaskByUuid(uuid string) (task Task, err error) {
	//获取当前时间
	currentTime := time.Now()
	// 获取当天的起始时间和结束时间
	startOfDay := currentTime.Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	// 根据uuid和时间范围查询任务记录
	result := Db.Where("uuid = ? AND created_at >= ? AND created_at < ?", uuid, startOfDay, endOfDay).First(&task)

	if result.Error != nil {
		// 如果记录不存在，创建新的任务记录
		task.CreatedAt = currentTime
		task.Sum = 0
		task.Uuid = uuid
		task.Day = int(currentTime.Weekday())
		result = Db.Create(&task)

		if result.Error != nil {
			err = result.Error
			return
		}
	}

	return
}

// DestroyOldByUuid 方法摧毁用户这周之前的所有任务数据
func DestroyOldByUuid(uuid string) {
	// 获取当前时间
	currentTime := time.Now()

	// 计算当前日期所在的周的起始日期（周一为一周的起始）
	startOfWeek := currentTime.AddDate(0, 0, -int(currentTime.Weekday())+1).Truncate(24 * time.Hour)

	// 根据uuid和时间范围删除任务记录
	Db.Where("uuid = ? AND created_at < ?", uuid, startOfWeek).Delete(&Task{})
}
