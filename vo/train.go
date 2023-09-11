package vo

import "studysystem/models"

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
