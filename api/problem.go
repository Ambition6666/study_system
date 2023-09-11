package api

import (
	"studysystem/internal/service/problem"
	"studysystem/vo"

	"github.com/gin-gonic/gin"
)

// 添加问题
func AddProblem(c *gin.Context) {
	v := new(vo.Add_problem_request)
	c.Bind(v)
	code, msg := problem.Add_problem(v)
	c.JSON(200, vo.Add_problem_response{
		Code: code,
		Msg:  msg,
	})
}

// 删除问题
func DeleteProblem(c *gin.Context) {
	v := new(vo.Delete_problem_resquest)
	c.Bind(v)
	code, msg := problem.Delete_problem(v.Qid)
	c.JSON(200, vo.Add_problem_response{
		Code: code,
		Msg:  msg,
	})
}
