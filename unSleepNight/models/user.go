package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Uuid     string `gorm:"size:64;not null;unique"`
	Name     string `gorm:"size:255;not null;unique"`
	Email    string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:64;not null"`
	Sex      string
}

// CreateSession 方法为临时user创建一个新的session
func (user *User) CreateSession() (session Session, err error) {
	//生成会话的Uuid
	uuid := createUuid()
	//使用gorm创建session记录
	session = Session{
		Uuid:   uuid,
		Email:  user.Email,
		UserId: user.ID,
	}
	result := Db.Create(&session)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// Session 方法用于获取现有用户的会话
func (user *User) Session() (session Session, err error) {
	//根据用户ID查询对应session
	err = Db.Where("user_id = ?", user.ID).First(&session).Error
	return
}

// Create 方法用于创建新用户并将用户信息保存到数据库中
func (user *User) Create() (err error) {
	//生成用户的UUID
	uuid := createUuid()

	//创建user记录
	user.Uuid = uuid
	result := Db.Create(user)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

// Delete 方法用于从数据库中删除用户
func (user *User) Delete() (err error) {
	//删除user记录
	result := Db.Delete(user)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

// Update 方法用于更新数据库中用户的信息
func (user *User) Update() (err error) {
	//更新user记录
	result := Db.Model(user).Updates(map[string]interface{}{"name": user.Name, "email": user.Email, "sex": user.Sex})
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// UserDeleteAll 方法用于删除数据库中的"所有"用户记录!
func UserDeleteAll() (err error) {
	//删除所有user记录
	result := Db.Delete(&User{})
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// Users 方法用于获取所有用户的信息
func Users() (users []User, err error) {
	//查询所有user用户记录
	result := Db.Find(&users)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// UserByEmail 方法用于根据邮箱查询用户记录
func UserByEmail(email string) (user User, err error) {
	//根据email查询用户记录
	result := Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

// UserByName 方法用于根据邮箱查询用户记录
func UserByName(name string) (user User, err error) {
	//根据email查询用户记录
	result := Db.Where("name = ?", name).First(&user)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

// UserByUuid 方法用于根据Uuid查询用户记录
func UserByUuid(uuid string) (user User, err error) {
	//根据Uuid来查询user记录
	result := Db.Where("uuid = ?", uuid).First(&user)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

// CreateThread 方法创建一个新的帖子
func (user *User) CreateThread(topic string) (thread Thread, err error) {
	//生成会话的Uuid
	uuid := createUuid()
	now := time.Now()
	thread = Thread{
		Uuid:      uuid,
		Topic:     topic,
		UserId:    user.ID,
		likes:     0,
		CreatedAt: now,
	}
	result := Db.Create(&thread)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// CreatePost 创建一个新的帖子回复
func (user *User) CreatePost(thread Thread, Body string) (post Post, err error) {
	uuid := createUuid()
	now := time.Now()
	post = Post{
		Uuid:      uuid,
		Body:      Body,
		UserId:    user.ID,
		ThreadId:  thread.ID,
		CreatedAt: now,
	}
	result := Db.Create(&post)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
