package vo

import (
	"encoding/json"
	clients "studysystem/api/proto"
	"studysystem/models"
)

type Video_list_resquest struct {
	MiniType  int   `form:"mini_type"`
	StageType int   `form:"stage_type"`
	Line_type int   `form:"line_type"`
	Offset    int64 `form:"offset"`
}
type Video_list_response struct {
	Code        int                `json:"code"`
	Videos      []models.Video     `json:"videos"`
	Study_route models.Study_route `json:"study_route"`
	MiniType    int                `json:"mini_type"`
	StageType   int                `json:"stage_type"`
}
type Problem_list_resquest struct {
	Vid uint `form:"vid"`
}
type Problem_list_response struct {
	Code int              `json:"code"`
	List []models.Problem `json:"list"`
}
type Commit_answer_resquest struct {
	Qid    uint    `form:"qid"`
	Answer []int32 `form:"answer"`
}
type Commit_answer_response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type Get_record_response struct {
	Code          int                 `json:"code"`
	Answer_record models.CommitRecord `json:"answer_record"`
}
type Code_answer struct {
	Msg   json.RawMessage `json:"msg"`
	MType string          `json:"type"`
}
type Commit_code struct {
	QID        int64  `json:"qid"`
	Code       string `json:"code"`
	LanguageID int64  `json:"language_id"`
}
type Hurt_beat struct {
	Msg string `json:"msg"`
}
type Get_problem struct {
	QID int64 `json:"qid"`
}
type Commit_response struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	MType int    `json:"type"`
}
type Commit_code_response struct {
	Code  int                  `json:"code"`
	Msg   *clients.JudgeResult `json:"msg"`
	MType int                  `json:"type"`
}
