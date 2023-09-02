package router

import (
	"studysystem/api"
	"studysystem/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/login", api.Login)       //登录
	r.POST("/register", api.Register) //注册
	idy := r.Group("/identified")
	idy.Use(middleware.Auth())
	{
		idy.GET("/userinfo", api.GetUserInfo)       //获取用户信息
		idy.POST("/updateinfo", api.UpdateUserInfo) //更新用户信息
		idy.GET("/avatar/", api.GetAvatar)          //获取用户头像
		idy.DELETE("/user", api.DeleteUser)         //删除用户
	}
	return r
}
