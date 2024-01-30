package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"unsleepNight/models"
)

// NewThread GET //thread/new
// 获取创建帖子的页面
func NewThread(c *gin.Context) {
	_, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	} else {
		generateHTML(c, nil, "layout", "auth.navbar", "new.thread")
	}
}

// CreateThread POST /thread/create
// 执行群组创建逻辑
func CreateThread(c *gin.Context) {
	sess, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login?message=请先登陆!")
		return
	}

	err = c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法解析的表单数据"})
		return
	}

	user, err := sess.User()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户数据查询错误"})
		return
	}

	//获取主题和标题
	topic := c.Request.PostFormValue("topic")
	title := c.Request.PostFormValue("title")
	if _, err := user.CreateThread(topic, title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建帖子失败"})
	}
	//看不懂的重定向
	c.Redirect(http.StatusFound, "/")
	return
}

//GET /thread/read
//通过ID渲染制定群组页面

func ReadThread(c *gin.Context) {

	//读取帖子
	uuid := c.Query("id")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不存在的帖子"})
		return
	}

	_, err = session(c)
	if err != nil {
		generateHTML(c, &thread, "layout", "navbar", "thread")
	} else {
		generateHTML(c, &thread, "layout", "auth.navbar", "auth.thread")
	}
}
