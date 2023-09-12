package vo

type Add_problem_request struct {
	VideoID     uint     `form:"video_id"`
	Description string   `form:"description"`
	Options     []string `form:"options"`
	Answer      []int32  `form:"answer"`
	ProblemType int      `form:"problem_type"`
	CodeID      int64    `form:"code_id"` //选择题id默认为0,如果是编程题,则它存的是问题id
}

// 跟删除操作响应共用
type Add_problem_response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type Delete_problem_resquest struct {
	Qid uint `form:"qid"`
}
