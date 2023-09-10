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

// 获取学习路线
func GetStudyRoute(c *gin.Context) {
	u := new(vo.Video_list_resquest)
	c.Bind(u)
	code, r, vlist := video.GetVideoList(u.Line_type, u.Type, u.Offset)
	c.JSON(200, vo.Video_list_response{
		Code:        code,
		Videos:      vlist,
		Study_route: r,
		Type:        u.Type,
	})
}
