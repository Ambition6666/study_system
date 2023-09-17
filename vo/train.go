package vo

import (
	"encoding/json"
	pri "studysystem/api/proto/private"

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
	Qid    uint    `form:"qid" json:"qid"`
	Answer []int32 `form:"answer" json:"answer"`
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

// 调试代码
type Debug_code struct {
	Input      string `json:"input"`
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
	Code  int          `json:"code"`
	Msg   *pri.Problem `json:"msg"`
	MType int          `json:"type"`
}

// 提交代码响应
type Commit_response struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	MType int    `json:"type"`
}

// 编程题是否做对响应
type Commit_code_response struct {
	Code  int              `json:"code"`
	Msg   *pri.JudgeResult `json:"msg"`
	MType int              `json:"type"`
}

// 调试代码响应
type debug_code_response struct {
	Code  int              `json:"code"`
	Msg   *pri.JudgeResult `json:"msg"`
	MType int              `json:"type"`
}

// 阶段测试请求
type Test_resquest struct {
	Line_type  int `form:"line_type"`
	State_type int `form:"state_type"`
}

// 测试响应
type Test_response struct {
	Code int              `form:"code"`
	Data []models.Problem `form:"data"`
}

// 提交测试
type Test_commit_resquest struct {
	Answer []json.RawMessage `json:"answer"`
}

// 测试结果
type Test_res struct {
	Istrue bool `json:"istrue"`
	QID    uint `json:"qid"`
}

type Test_res_s struct {
	Res   []Test_res `json:"res"`
	Score int        `json:"score"`
}

// 提交测试响应
type Test_commit_response struct {
	Code int        `json:"code"`
	Data Test_res_s `json:"data"`
}
