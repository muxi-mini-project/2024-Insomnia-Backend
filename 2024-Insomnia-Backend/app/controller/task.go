package controller

import (
	. "Insomnia/app/core/helper"
	"Insomnia/app/response"
	"Insomnia/app/service"
	"github.com/gin-gonic/gin"
)

type Task struct{}

var taskService *service.TaskService

func (t *Task) GetTask(c *gin.Context) {
	//获取当前用户的Uuid
	Uuid, _ := c.Get("Uuid")
	uuid := Uuid.(string)
	err, task := taskService.GetTask(uuid)
	if err != nil {
		response.FailMsgData(c, err.Error(), response.GetTaskResponse{})
		return
	}
	todayTask := response.GetTaskResponse{Sum: task.Sum, Day: uint(task.Day)}
	response.OkMsgData(c, "获取今日数据成功", todayTask)
	return
}

func (t *Task) UpTask(c *gin.Context) {
	//获取当前用户的Uuid
	Uuid, _ := c.Get("Uuid")
	uuid := Uuid.(string)
	err := taskService.UpTask(uuid)
	if err != nil {
		response.FailMsg(c, err.Error())
		return
	}
	response.FailMsg(c, "更新数据成功!")
	return
}

func (t *Task) GetAllTask(c *gin.Context) {
	//获取当前用户的Uuid
	Uuid, _ := c.Get("Uuid")
	uuid := Uuid.(string)
	err, allTask := taskService.GetAllTask(uuid)
	if err != nil {
		Danger(err, "获取本周数据失败")
		response.FailMsgData(c, err.Error(), []response.GetAllTaskResponse{})
		return
	}
	response.OkMsgData(c, "获取本周数据成功", allTask)
	return
}
