package api

import (
	"fmt"
	"studysystem/config"
	"studysystem/internal/repository"
	"studysystem/internal/service/user"
	"studysystem/vo"

	"github.com/gin-gonic/gin"
)

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	id := GetUserID(c)
	uinfo := user.GetUserInfo(id)
	c.JSON(200, *uinfo)
}

// 修改用户信息
func UpdateUserInfo(c *gin.Context) {
	id := GetUserID(c)
	action := c.PostForm("action")
	if action == "2" {
		file, err := c.FormFile("data")
		if err != nil {
			c.JSON(500, "传输失败")
		}
		c.SaveUploadedFile(file, config.LocalPath+file.Filename)
		code, data := user.UpdateUserInfo(id, 2, config.LocalPath+file.Filename)
		c.JSON(code, data)
		return
	}
	u := new(vo.Update_user_request)
	c.Bind(u)

	code, data := user.UpdateUserInfo(id, u.Action, u.Data)
	c.JSON(code, data)
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id := GetUserID(c)
	code, data := user.DeleteUser(id)
	c.JSON(code, data)
}

// 获取用户头像
func GetAvatar(c *gin.Context) {
	id := GetUserID(c)
	data, err := repository.Get_local_avatar_path(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, "获取头像失败")
	}
	c.File(data)
}