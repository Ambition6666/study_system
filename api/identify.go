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
	code, msg := login.IdentifyCode(user_info.Email, user_info.Code)
	if code != 200 {
		c.JSON(code, vo.Register_response{
			Code: code,
			Msg:  msg,
		})
		return
	}
	code, msg = login.Register(user_info.Email, user_info.Pwd)
	c.JSON(code, vo.Register_response{
		Code: code,
		Msg:  msg,
	})
}

// 用户密码登录
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

// 用户验证码登录
func Login_by_auth_code(c *gin.Context) {
	user_info := new(vo.Login_by_auth_code_request)
	c.Bind(user_info)
	code, msg := login.Login_by_auth_code(user_info.Email, user_info.AuthCode)
	if code != 200 {
		c.JSON(200, vo.Login_response{
			Code:  code,
			Msg:   msg,
			Token: "",
		})
	} else {
		c.JSON(200, vo.Login_response{
			Code:  code,
			Msg:   "登录成功",
			Token: msg,
		})
	}
}

// 获取验证码
func GetAuthCode(c *gin.Context) {
	em := c.Query("email")
	login.SendAuthCode(em)
	c.JSON(200, "发送成功")
}

// 获取user的id
func GetUserID(c *gin.Context) int64 {
	id := c.GetInt64("id")
	return id
}
