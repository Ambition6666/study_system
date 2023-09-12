package api

import (
	"studysystem/internal/service/train"
	"studysystem/internal/service/video"
	websokcet "studysystem/internal/service/websocket"
	"studysystem/vo"

	"github.com/gin-gonic/gin"
)

// 获取学习路线
func GetStudyRoute(c *gin.Context) {
	u := new(vo.Video_list_resquest)
	c.Bind(u)
	code, r, vlist := video.GetVideoList(u.Line_type, u.MiniType, u.StageType)
	c.JSON(200, vo.Video_list_response{
		Code:        code,
		Videos:      vlist,
		Study_route: r,
		MiniType:    u.MiniType,
		StageType:   u.StageType,
	})
}

// 获取问题
func GetProblems(c *gin.Context) {
	v := new(vo.Problem_list_resquest)
	c.Bind(v)
	code, plist := train.Get_problem_list(v.Vid)
	c.JSON(200, vo.Problem_list_response{
		Code: code,
		List: plist,
	})
}

// 回答问题
func CommitAnswer(c *gin.Context) {
	uid := GetUserID(c)
	v := new(vo.Commit_answer_resquest)
	c.Bind(v)
	code, msg := train.CommitAnswer(v.Answer, v.Qid, uid)
	c.JSON(200, vo.Commit_answer_response{
		Code: code,
		Msg:  msg,
	})
}

// 回答编程问题
func CommitCodeAnswer(c *gin.Context) {
	cc := websokcet.Up(c)
	cl := websokcet.NewClient(c.Query("uid")+c.Query("qid"), cc.RemoteAddr().String(), cc)
	go cl.Read()
	go cl.Write()
	go cl.TimeOutClose()
	websokcet.Manager.Register <- cl
}

// 获取记录
func GetRecord(c *gin.Context) {
	qid := c.Query("qid")
	code, r := train.GetRecord(qid, GetUserID(c))
	c.JSON(200, vo.Get_record_response{
		Code:          code,
		Answer_record: *r,
	})
}
