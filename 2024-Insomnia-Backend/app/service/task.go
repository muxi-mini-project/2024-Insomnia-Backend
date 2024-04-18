package service

import (
	"Insomnia/app/models"
	. "Insomnia/app/models"
	"Insomnia/app/response"
	"fmt"
)

type TaskService struct{}

// UpTask /POST /task1/upTask
// 更新任务数据
func (t *TaskService) UpTask(uuid string) (err error) {

	//获取当月任务数据
	task, err := GetTaskByUuid(uuid)
	if err != nil {
		return fmt.Errorf("用户任务数据查询错误:%v", err)
	}

	//更新任务数据
	err = task.Update()
	if err != nil {
		return fmt.Errorf("用户任务数据更新失败%v", err)
	}

	return nil
}

// GetTask /GET /task1/getTask
// 获取本日任务数据
func (t *TaskService) GetTask(uuid string) (error, models.Task) {

	//删除用户所有这周之前的任务数据
	DestroyOldByUuid(uuid)

	//获取本日的任务数据
	task, err := GetTaskByUuid(uuid)
	if err != nil {
		return fmt.Errorf("用户任务数据查询错误%v", err), models.Task{}
	}

	return nil, task
}

// GetAllTask /GET /task1/getAllTask
// 获取本周的数据
func (t *TaskService) GetAllTask(uuid string) (error, []response.GetAllTaskResponse) {
	//删除用户所有这周之前的任务数据
	DestroyOldByUuid(uuid)
	var allTask []response.GetAllTaskResponse
	var t4 response.GetAllTaskResponse
	var t1 response.GetTaskResponse
	var t2 []response.GetTaskResponse
	//获取本周的任务数据
	task, err := GetAllTaskByUuid(uuid)
	if err != nil {
		return fmt.Errorf("用户任务数据查询错误:%v", err), []response.GetAllTaskResponse{}
	}
	for _, t3 := range task {
		t1 = response.GetTaskResponse{
			Sum: t3.Sum,
			Day: uint(t3.Day),
		}
		t2 = append(t2, t1)
		t4 = response.GetAllTaskResponse{AllTask: t2}
		allTask = append(allTask, t4)
	}
	return nil, allTask
}
