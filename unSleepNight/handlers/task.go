package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpTask /POST /task/upTask
// 更新任务数据
func UpTask(c *gin.Context) {
	sess, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login?message=请先登陆!")
		return
	}

	user, err := sess.User()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户数据查询错误"})
		return
	}

	//获取当月任务数据
	task, err := user.GetTaskByUuid()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户任务数据查询错误"})
		return
	}

	//更新任务数据
	err = task.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户任务数据更新失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户任务数据更新成功"})
	return
}

// GetTask /GET /task/getTask
// 获取当月任务数据
func GetTask(c *gin.Context) {
	sess, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login?message=请先登陆!")
		return
	}

	user, err := sess.User()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户数据查询错误"})
		return
	}

	//删除用户所有去年的任务数据
	user.DestroyOld()

	//获取当月的任务数据
	task, err := user.GetTaskByUuid()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户该月数据查询错误"})
		return
	}

	//返回task值
	c.JSON(http.StatusOK, task)
	return
}

// GetAllTask /GET /task/getAllTask
// 获取整年的数据
func GetAllTask(c *gin.Context) {
	sess, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login?message=请先登陆!")
		return
	}

	user, err := sess.User()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户数据查询错误"})
		return
	}

	//删除用户所有去年的任务数据
	user.DestroyOld()

	task, err := user.GetAllTaskByUuid()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户全年数据查询错误"})
		danger(err, "用户全年数据查询错误")
		return
	}

	//返回task值
	c.JSON(http.StatusOK, task)
	return
}
