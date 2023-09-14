package vo

import (
	"encoding/json"
	judge "studysystem/api/proto/judge"
	problemrpc "studysystem/api/proto/problem"
	"studysystem/models"
)

// 学习路线获取视频请求
type Video_list_resquest struct {
	MiniType  int   `form:"mini_type"`
	StageType int   `form:"stage_type"`
	Line_type int   `form:"line_type"`
	Offset    int64 `form:"offset"`
}

// 学习路线获取视频响应
type Video_list_response struct {
	Code        int                `json:"code"`
	Videos      []models.Video     `json:"videos"`
	Study_route models.Study_route `json:"study_route"`
	MiniType    int                `json:"mini_type"`
	StageType   int                `json:"stage_type"`
}

// 视频播放完后获取问题请求
type Problem_list_resquest struct {
	Vid uint `form:"vid"`
}

// 视频播放完后获取问题响应
type Problem_list_response struct {
	Code int              `json:"code"`
	List []models.Problem `json:"list"`
}

// 选择题提交答案请求
type Commit_answer_resquest struct {
	Qid    uint    `form:"qid"`
	Answer []int32 `form:"answer"`
}

// 选择题提交答案是否成功响应
type Commit_answer_response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 获取选择题回答是否正确响应
type Get_record_response struct {
	Code          int                 `json:"code"`
	Answer_record models.CommitRecord `json:"answer_record"`
}

// 编程题开websocket请求
type Code_answer struct {
	Msg   json.RawMessage `json:"msg"`
	MType string          `json:"type"`
}

// 提交做题代码
type Commit_code struct {
	UID        uint   `json:"uid"`
	QID        uint   `json:"qid"`
	Code       string `json:"code"`
	LanguageID int64  `json:"language_id"`
}

// 心跳请求
type Hurt_beat struct {
	Msg string `json:"msg"`
}

// 获取编程问题内容
type Get_problem struct {
	QID uint `json:"qid"`
}

// 获取编程问题响应
type Get_problem_response struct {
	Code  int                 `json:"code"`
	Msg   *problemrpc.Problem `json:"msg"`
	MType int                 `json:"type"`
}

// 提交代码响应
type Commit_response struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	MType int    `json:"type"`
}

// 编程题是否做对响应
type Commit_code_response struct {
	Code  int                `json:"code"`
	Msg   *judge.JudgeResult `json:"msg"`
	MType int                `json:"type"`
}
