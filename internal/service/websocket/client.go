package websokcet

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"studysystem/clients"
	"studysystem/internal/repository"
	"studysystem/logs"
	"studysystem/models"
	"studysystem/vo"
	"time"

	rpc "studysystem/clients"

	pri "studysystem/api/proto/private"

	web "github.com/gorilla/websocket"
)

type Client struct {
	Name       string
	Addr       string
	Send       chan []byte
	Conn       *web.Conn
	Login_time time.Time
}

func NewClient(n string, addr string, cc *web.Conn) *Client {
	return &Client{
		Name:       n,
		Addr:       addr,
		Send:       make(chan []byte, 100),
		Conn:       cc,
		Login_time: time.Now(),
	}
}

// 发送给客户端
func (c *Client) Write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)

		}
	}()

	defer func() {
		Manager.Unregister <- c
		c.Conn.Close()
		fmt.Println("Client发送数据 defer", c)
	}()

	for v := range c.Send {
		err := c.Conn.WriteMessage(web.TextMessage, v)
		if err != nil {
			Manager.Unregister <- c
			c.Conn.Close()
			fmt.Println("Client发送数据 defer", c)
		}
	}
}

// 写数据给客户端
func (c *Client) SendMessage(msg []byte) {
	if c == nil {

		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("SendMsg stop:", r, string(debug.Stack()))
		}
	}()

	c.Send <- msg
}

// 接收客户端发来的数据
func (c *Client) Read() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("write stop", string(debug.Stack()), e)
		}
	}()
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg)
		ParseDate(c, msg)
	}
}
func ParseDate(cl *Client, msg []byte) {
	v := new(vo.Code_answer)
	err := json.Unmarshal(msg, v)
	if err != nil {
		fmt.Println("整体消息:", err)
		return
	}
	ctx := context.Background()
	switch v.MType {
	case "1":
		GetCodeProblem(cl, ctx, v)
	case "2":
		CommitCodeAnswer(cl, ctx, v)
	case "3":
		Manager.IsLogin <- cl
	default:
		return
	}
}

// 定时检测客户端是否连接超时
func (c *Client) TimeOutClose() {
	for {
		time.Sleep(10 * time.Second)
		t := time.Now()
		if t.After(c.Login_time.Add(60 * time.Second)) {
			Manager.Unregister <- c
			c.Conn.Close()
			return
		}
	}
}

// 提交题目
func CommitCodeAnswer(cl *Client, ctx context.Context, v *vo.Code_answer) {
	val := new(vo.Commit_code)
	err := json.Unmarshal(v.Msg, val)
	if err != nil {
		logs.SugarLogger.Debugln("解析消息错误", err)
		return
	}
	q := repository.Get_problem(val.QID)
	res, err := rpc.ProCli.Judge(ctx, &pri.JudgeRequest{
		ProblemID: q.CodeID,
		Code:      []byte(val.Code),
		LangID:    val.LanguageID,
	})
	if err != nil || res.StatusCode != 1000 {
		logs.SugarLogger.Debugln("提交题目错误", err)
		v := vo.Commit_response{
			Code:  500,
			Msg:   "提交失败",
			MType: 2,
		}
		msg, _ := json.Marshal(v)
		cl.SendMessage(msg)
		return
	} else {
		v := vo.Commit_response{
			Code:  200,
			Msg:   res.JudgeID,
			MType: 2,
		}
		msg, _ := json.Marshal(v)
		cl.SendMessage(msg)
	}
	res1, err := rpc.ProCli.GetResult(ctx, &pri.GetResultRequest{
		JudgeID: res.JudgeID,
	})
	if err != nil || res1.StatusCode != 1000 {
		logs.SugarLogger.Debugln("获取结果:", err)
		v := vo.Commit_code_response{
			Code:  500,
			Msg:   res1.Result,
			MType: 2,
		}
		msg, _ := json.Marshal(v)
		cl.SendMessage(msg)
		return
	} else {
		v := vo.Commit_code_response{
			Code:  200,
			Msg:   res1.Result,
			MType: 2,
		}
		msg, _ := json.Marshal(v)
		cl.SendMessage(msg)
		record := new(models.CommitRecord)
		record.Qid = q.ID
		if res1.Result.Status == 10 {
			record.IsTrue = true
		} else {
			record.IsTrue = false
		}
		record.Uid = int64(val.UID)
		repository.CreateCommitRecord(record)
	}
}
func GetCodeProblem(cl *Client, ctx context.Context, v *vo.Code_answer) {
	val := new(vo.Get_problem)
	err := json.Unmarshal(v.Msg, val)
	if err != nil {
		logs.SugarLogger.Debugln("解析消息错误", err)
		return
	}
	q := repository.Get_problem(val.QID)
	res, err := clients.ProCli.GetProblem(context.Background(), &pri.GetProblemRequest{
		ProblemID: q.CodeID,
	})
	if err != nil {
		logs.SugarLogger.Debugln("获取问题错误", err)
	}
	if res.StatusCode != 1000 {
		v := vo.Get_problem_response{
			Code:  500,
			Msg:   res.Problem,
			MType: 1,
		}
		msg, _ := json.Marshal(v)
		cl.SendMessage(msg)
		return
	} else {
		v := vo.Get_problem_response{
			Code:  200,
			Msg:   res.Problem,
			MType: 1,
		}
		msg, _ := json.Marshal(v)
		cl.SendMessage(msg)
	}

}
