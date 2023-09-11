package api

import (
	studyroute "studysystem/internal/service/study_route"
	"studysystem/vo"

	"github.com/gin-gonic/gin"
)

func Add_study_route(c *gin.Context) {
	v := new(vo.Add_study_route_resquest)
	c.Bind(v)
	code, msg := studyroute.Add_study_route(v.Description, v.Line_type)
	c.JSON(200, vo.Add_study_route_response{
		Code: code,
		Msg:  msg,
	})
}
