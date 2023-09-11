package vo

type Add_problem_request struct {
	VideoID     uint     `form:"video_id"`
	Description string   `form:"description"`
	Options     []string `form:"options"`
	Answer      []int32  `form:"answer"`
	ProblemType int      `form:"problem_type"`
}

// 跟删除操作响应共用
type Add_problem_response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type Delete_problem_resquest struct {
	Qid uint `form:"qid"`
}
