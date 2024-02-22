package controller

import (
	"Insomnia/app/common/tool"
	. "Insomnia/app/core/helper"
	"Insomnia/app/request"
	"Insomnia/app/response"
	"Insomnia/app/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Auth struct{}

var authService *service.AuthService

// Login 用户登录
// @Summary 用户登录接口
// @Description 用户登录接口
// @Tags Auth
// @Accept json
// @Produce json
// @Param email body string true "邮箱"
// @Param password body string true "密码"
// @Success 200 {object} LoginResponse "登录成功"
// @Failure 400 {object} string "请求参数错误"
// @Failure 500 {object} string "内部错误"
func (a *Auth) Login(c *gin.Context) {
	//定义一个Login请求类型的结构体
	req := &request.LoginReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}

	//调用服务层来获取一个token
	token, err := authService.Login(req.Email, tool.Encrypt(req.Password))
	if err != nil {
		Danger(err, "获取token失败")
		response.FailMsg(c, fmt.Sprintf("获取token失败: %v", err))
		return
	}

	//返回消息捏
	response.OkMsgData(c, "登录成功", response.LoginResponse{Token: token})
}

// Signup 用户注册
// @Summary 用户注册接口
// @Description 用户注册接口
// @Tags Auth
// @Accept json
// @Produce json
// @Param email body string true "邮箱"
// @Param password body string true "密码"
// @Param verificationCode body string true "验证码"
// @Param sex body int true "性别"
// @Success 200 {object} string "注册成功"
// @Failure 400 {object} string "请求参数错误"
// @Failure 500 {object} string "内部错误"
func (a *Auth) Signup(c *gin.Context) {
	//定义一个Login请求类型的结构体
	sur := &request.SignupReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(sur); err != nil {
		Danger(err, "无法解析的表单")
		response.FailMsg(c, fmt.Sprintf("无法解析: %v", err))
		return
	}

	//调用服务层来注册一个账户
	err := authService.Signup(*sur)
	if err != nil {
		Danger(err, "注册时服务器发生错误")
		response.FailMsg(c, fmt.Sprintf("注册时服务器发生错误: %v", err))
		return
	}

	//返回消息捏
	response.OkMsg(c, "注册成功!")
}

// ChangePassword 更改密码
// @Summary 更改密码接口
// @Description 更改密码接口
// @Tags Auth
// @Accept json
// @Produce json
// @Param email body string true "邮箱"
// @Param verificationCode body string true "验证码"
// @Param newPassword body string true "新密码"
// @Success 200 {object} string "密码更改成功"
// @Failure 400 {object} string "请求参数错误"
// @Failure 500 {object} string "内部错误"
func (a *Auth) ChangePassword(c *gin.Context) {
	//定义一个Login请求类型的结构体
	cp := &request.ChangePasswordReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(cp); err != nil {
		Danger(err, "无法解析的表单")
		response.FailMsg(c, fmt.Sprintf("无法解析: %v", err))
		return
	}

	//调用服务层来更新密码
	err := authService.ChangePassword(*cp)
	if err != nil {
		Danger(err, "更新密码失败")
		response.FailMsg(c, fmt.Sprintf("更新密码失败: %v", err))
		return
	}

	//返回消息捏
	response.OkMsg(c, "更改密码成功!")
}

// ChangeAvatar 更改头像
// @Summary 更改头像接口
// @Description 更改头像接口
// @Tags Auth
// @Accept json
// @Produce json
// @Param sex body int true "头像"
// @Param Uuid header string true "用户唯一标识" default(uuid)
// @Success 200 {object} string "头像更改成功"
// @Failure 400 {object} string "请求参数错误"
// @Failure 500 {object} string "内部错误"
func (a *Auth) ChangeAvatar(c *gin.Context) {
	//定义一个Login请求类型的结构体
	cs := &request.ChangeAvatarReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(cs); err != nil {
		Danger(err, "无法解析的表单")
		response.FailMsg(c, fmt.Sprintf("无法解析: %v", err))
		return
	}

	Uuid, _ := c.Get("Uuid")
	uuid := Uuid.(string)
	//调用服务层来更新头像
	err := authService.ChangeAvatar(*cs, uuid)
	if err != nil {
		Danger(err, "更新头像失败")
		response.FailMsg(c, fmt.Sprintf("更新头像失败: %v", err))
		return
	}

	//返回消息捏
	response.OkMsg(c, "更改头像成功!")
}

// GetMyData 获取数据
// @Summary 获取数据
// @Description 获取数据
// @Tags Auth
// @Accept json
// @Produce json
// @Param Uuid header string true "用户唯一标识"
// @Success 200 {object} GetMyDataResponse "获取数据成功"
// @Failure 400 {object} string "请求参数错误"
// @Failure 500 {object} string "内部错误"
func (a *Auth) GetMyData(c *gin.Context) {
	Uuid, _ := c.Get("Uuid")
	uuid := Uuid.(string)
	myData, err := authService.GetMyData(uuid)
	if err != nil {
		Danger(err, "获取用户信息失败")
		response.FailMsgData(c, fmt.Sprintf("获取用户信息失败: %v", err), myData)
		return
	}
	response.OkMsgData(c, "获取用户信息成功", myData)
	return
}

//func (a *Auth) MyMessage(c *gin.Context) {
//	Uuid, _ := c.Get("Uuid")
//	uuid := Uuid.(string)
//	user, err := authService.MyMessage(uuid)
//	if err != nil {
//		Danger(err, "获取用户信息失败")
//		response.FailMsgData(c, fmt.Sprintf("获取用户信息失败: %v", err), user)
//		return
//	}
//	response.OkMsgData(c, "获取用户信息成功", user)
//	return
//}
