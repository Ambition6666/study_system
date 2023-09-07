package api

import (
	"studysystem/internal/service/video"
	"studysystem/vo"

	"github.com/gin-gonic/gin"
)

func AddVideo(c *gin.Context) {
	u := new(vo.Add_video_resquest)
	c.Bind(u)
	code, msg := video.AddVideo(u)
	c.JSON(200, vo.Add_video_response{
		Msg:  msg,
		Code: code,
	})
}
