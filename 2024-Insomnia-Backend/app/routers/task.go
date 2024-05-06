package routers

import "Insomnia/app/core/middlewares"

// useTask 获取当日数据的路由
func (r *router) useTask() {
	taskRouter := r.Group("task1")
	taskRouter.POST("/getTask", middlewares.UseJwt(), r.task.GetTask)
	taskRouter.POST("/upTask", middlewares.UseJwt(), r.task.UpTask)
	taskRouter.POST("/getAllTask", middlewares.UseJwt(), r.task.GetAllTask)
}
