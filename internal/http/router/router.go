package router

import (
	"studysystem/api"
	"studysystem/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	a := r.Group("/api")
	a.POST("/login", api.Login)                    //登录
	a.POST("/loginbycode", api.Login_by_auth_code) //登录
	a.POST("/register", api.Register)              //注册
	a.GET("/authcode", api.GetAuthCode)            //获取验证码
	idy := a.Group("/identified")
	idy.Use(middleware.Auth())
	{
		idy.GET("/userinfo", api.GetUserInfo)       //获取用户信息
		idy.POST("/updateinfo", api.UpdateUserInfo) //更新用户信息
		idy.GET("/avatar/", api.GetAvatar)          //获取用户头像
		idy.DELETE("/user", api.DeleteUser)         //删除用户
	}
	return r
}
