package api

import (
	"studysystem/internal/service/video"
	"studysystem/vo"

	"github.com/gin-gonic/gin"
)

// 添加视频
func AddVideo(c *gin.Context) {
	u := new(vo.Add_video_resquest)
	c.Bind(u)
	code, msg := video.AddVideo(u)
	c.JSON(200, vo.Add_video_response{
		Msg:  msg,
		Code: code,
	})
}

// 删除视频
func DeleteVideo(c *gin.Context) {
	u := new(vo.Delete_video_resquest)
	c.Bind(u)
	code, msg := video.DeleteVideo(u.Vid)
	c.JSON(200, vo.Add_video_response{
		Msg:  msg,
		Code: code,
	})
}
