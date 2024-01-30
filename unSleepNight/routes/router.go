package routes

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	//创建gin.Engine路由器示例
	router := gin.Default()
	for _, route := range webRoutes {
		//将每一个Gin路由应用到路由器
		switch route.Method {
		case "GET":
			router.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			router.POST(route.Pattern, route.HandlerFunc)
		default:
			// 默认使用 Any 方法处理所有HTTP方法(什么鸟方法)
			router.Any(route.Pattern, route.HandlerFunc)
		}
	}
	return router
}
