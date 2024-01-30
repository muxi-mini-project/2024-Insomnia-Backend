package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"unsleepNight/models"
)

// GET /login
// 登陆页面

func Login(c *gin.Context) {
	t := parseTemplateFiles("auth.layout", "navbar", "login")
	err := t.Execute(c.Writer, nil)
	if err != nil {
		return
	}
}

// GET /signup
// 注册页面

func Signup(c *gin.Context) {
	generateHTML(c, nil, "auth.layout", "navbar", "signup")
}

// POST /sign
//注册新用户

func SignupAccount(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法解析的表单数据"})
		return
	}
	user := models.User{
		Name:     c.Request.PostFormValue("name"),
		Email:    c.Request.PostFormValue("email"),
		Password: c.Request.PostFormValue("password"),
		Sex:      c.Request.PostFormValue("sex"),
	}

	user, err = models.UserByEmail(c.Request.PostFormValue("email"))
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该邮箱已被注册"})
		return
	}

	user, err = models.UserByEmail(c.Request.PostFormValue("name"))
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该用户名已被注册"})
		return
	}

	if err := user.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "注册失败"})
	}

	c.Redirect(http.StatusFound, "/signup")
}

//POST /authenticate
//通过邮箱和密码字段对用户进行登陆认证

func Authenticate(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法解析的表单数据"})
		return
	}
	user, err := models.UserByEmail(c.Request.PostFormValue("email"))
	if err := user.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该邮箱还未注册"})
		return
	}
	if user.Password == models.Encrypt(c.Request.PostFormValue("password")) {
		session, err := user.CreateSession()
		//缺少错误返回
		if err != nil {

			return
		}
		cookie := http.Cookie{
			Name:     "cookie",
			Value:    session.Uuid,
			HttpOnly: true,
			Expires:  time.Now().Add(24 * 30 * time.Hour),
			Path:     "/",
		}
		http.SetCookie(c.Writer, &cookie)
		c.Redirect(http.StatusFound, "/")
	} else {
		c.Redirect(http.StatusFound, "/login")
	}

}

//GET /logout
//用户退出

func Logout(c *gin.Context) {
	//判断是否登陆
	cookie, err := c.Request.Cookie("cookie")
	if err != http.ErrNoCookie {
		session := models.Session{Uuid: cookie.Value}

		err := session.DeleteByUuid()
		if err != nil {
			// 错误处理，例如记录日志或向用户提供适当的错误信息
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
			return
		}

		// 将名为 "user" 的 Cookie 的 MaxAge 设置为负值，达到清除的效果
		c.SetCookie("cookie", "", -1, "/", "localhost", false, true)
	}
	c.Redirect(http.StatusFound, "/")
}
