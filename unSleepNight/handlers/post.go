package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"unsleepNight/models"
)

// PostThread POST /thread/post
// 在指定群组下创建新主题
func PostThread(c *gin.Context) {
	sess, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	err = c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法解析的表单数据"})
		return
	}
	user, err := sess.User()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "不存在的会话"})
		return
	}
	body := c.Request.PostFormValue("body")
	uuid := c.Request.PostFormValue("uuid")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不存在的帖子"})
		return
	}
	if _, err := user.CreatePost(thread, body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建回复"})
	}
	url := fmt.Sprintf("/thread/read?id=", uuid)
	c.Redirect(http.StatusFound, url)
	return
}
