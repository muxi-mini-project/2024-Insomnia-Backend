package controller

import (
	. "Insomnia/app/common/tool"
	. "Insomnia/app/core/config"
	. "Insomnia/app/core/helper"
	"Insomnia/app/models"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"log"
	"net/http"
	"net/smtp"
)

type SendEmail struct{}

// SendEmail 发送邮件验证码
func (e *SendEmail) SendEmail(c *gin.Context) {
	//设置log参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	//获取全局配置
	config := LoadConfig()

	em := email.NewEmail()

	//设置发送方的邮箱,此处可以写自己的邮箱
	em.From = config.Email.UserName + "<" + config.Email.Sender + ">"

	//获取随机验证码
	random := GetRandom()

	//创造一个临时的CheckEmail
	Email := models.CheckEmail{
		Email:            c.Request.PostFormValue("email"),
		VerificationCode: random,
	}

	err := Email.CreateRedis()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "存储验证码失败"})
		Danger(err, "存储验证码到Redis失败")
		return
	}
	//设置接收方的邮箱
	em.To = []string{Email.Email}

	// 设置主题
	em.Subject = "验证码"

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte(Email.VerificationCode + "(验证码将在5分钟后失效,请不要告诉其他人，并尽快注册。")

	//设置服务器相关的配置
	err = em.Send(config.Email.Smtp, smtp.PlainAuth("", config.Email.Sender, config.Email.Password, "smtp.qq.com"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "邮箱服务器出错"})
		Danger(err, "邮箱服务器配置失败")
		return
	}

	//提示发送成功
	c.JSON(http.StatusOK, gin.H{"message": "发送验证码成功!"})
	return
}
