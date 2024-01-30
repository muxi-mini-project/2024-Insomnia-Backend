package routes

import (
	"github.com/gin-gonic/gin"
	"unsleepNight/handlers"
)

// GinRoute 存储单个的当前路由
type GinRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

// GinRoutes 定义一个可以存放所有类型路由的切片
type GinRoutes []GinRoute

// 定义所有Gin路由
var webRoutes = GinRoutes{
	{
		"home",
		"GET",
		"/",
		handlers.Index,
	},
	{
		"signup",
		"GET",
		"/signup",
		handlers.Signup,
	},
	{
		"signupAccount",
		"POST",
		"/signup_account",
		handlers.SignupAccount,
	},
	{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},
	{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},
	{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
}
