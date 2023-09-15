package api

import (
	"path"
	"path/filepath"
	"strconv"
	"studysystem/config"
	"studysystem/internal/service/user"
	"studysystem/logs"
	"studysystem/vo"

	"github.com/gin-gonic/gin"
)

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	id := GetUserID(c)
	uinfo := user.GetUserInfo(id)
	c.JSON(200, vo.Get_user_info_response{
		Code:      200,
		User_info: *uinfo,
	})
}

// 修改用户信息
func UpdateUserInfo(c *gin.Context) {
	id := GetUserID(c)
	action := c.PostForm("action")
	if action == "2" {
		file, err := c.FormFile("data")
		if err != nil {
			c.JSON(200, vo.Update_user_response{
				Code: 500,
				Data: "传输失败",
			})
			logs.SugarLogger.Errorf("传输失败:%v", err)
			return
		}
		lname := path.Ext(file.Filename)
		file.Filename = strconv.FormatInt(id, 10) + lname
		err = c.SaveUploadedFile(file, filepath.Join(config.LocalPath, "avatar", file.Filename))
		if err != nil {
			c.JSON(200, vo.Update_user_response{
				Code: 500,
				Data: "上传失败",
			})
			logs.SugarLogger.Errorf("上传失败:%v", err)
		}
		code, data := user.UpdateUserInfo(id, 2, file.Filename)
		c.JSON(200, vo.Update_user_response{
			Code: code,
			Data: data,
		})
		return
	}
	u := new(vo.Update_user_request)
	c.Bind(u)

	code, data := user.UpdateUserInfo(id, u.Action, u.Data)
	c.JSON(200, vo.Update_user_response{
		Code: code,
		Data: data,
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id := GetUserID(c)
	code, data := user.DeleteUser(id)
	c.JSON(200, vo.Delete_user_response{
		Code: code,
		Data: data,
	})
}
