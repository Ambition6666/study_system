package router

import (
	"studysystem/api"
	"studysystem/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors()) //跨域验证
	a := r.Group("/api")
	a.POST("/login", api.Login)                    //登录
	a.POST("/loginbycode", api.Login_by_auth_code) //登录
	a.POST("/register", api.Register)              //注册
	a.GET("/authcode", api.GetAuthCode)            //获取验证码
	a.GET("/answerCode", api.CommitCodeAnswer)     //题目
	idy := a.Group("/identified")
	idy.Use(middleware.Auth())
	{
		idy.GET("/userinfo", api.GetUserInfo)       //获取用户信息
		idy.POST("/updateinfo", api.UpdateUserInfo) //更新用户信息
		idy.GET("/avatar/", api.GetAvatar)          //获取用户头像
		idy.DELETE("/user", api.DeleteUser)         //删除用户
		sty := idy.Group("/study")
		{
			sty.POST("/getstudyroute", api.GetStudyRoute) //获取学习路线
			sty.POST("/getquestion", api.GetProblems)     //获取题目
			sty.POST("/answer", api.CommitAnswer)         //回答选择题
			sty.GET("/record", api.GetRecord)             //得到答题记录
		}
	}
	admin := a.Group("/admin") //管理员操作
	admin.Use(middleware.Auth(), middleware.If_admin())
	{
		admin.POST("/addvideo", api.AddVideo)             //管理员添加视频
		admin.DELETE("/deletevideo", api.DeleteVideo)     //管理员删除视频
		admin.POST("/addproblem", api.AddProblem)         //管理员添加问题
		admin.DELETE("/deleteproblem", api.DeleteProblem) //管理员删除问题
		admin.POST("/addstudyroute", api.Add_study_route) //管理员添加学习路线
	}
	return r
}
