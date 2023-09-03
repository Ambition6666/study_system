package api

import (
	"studysystem/internal/service/login"
	"studysystem/vo"

	"github.com/gin-gonic/gin"
)

// 用户注册
func Register(c *gin.Context) {
	user_info := new(vo.Register_resquest)
	c.Bind(user_info)
	code, msg := login.Register(user_info.Email, user_info.Pwd)
	c.JSON(code, vo.Register_response{
		Code: code,
		Msg:  msg,
	})
}

// 用户登录
func Login(c *gin.Context) {
	user_info := new(vo.Login_resquest)
	c.Bind(user_info)
	code, data := login.Login(user_info.Email, user_info.Pwd)
	if code != 200 {
		c.JSON(200, vo.Login_response{
			Code:  code,
			Msg:   data,
			Token: "",
		})
	} else {
		c.JSON(200, vo.Login_response{
			Code:  code,
			Msg:   "登录成功",
			Token: data,
		})
	}
}

// 获取user的id
func GetUserID(c *gin.Context) int64 {
	id, ok := c.Get("id")
	if !ok {
		return 0
	}
	return id.(int64)
}
