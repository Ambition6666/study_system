package vo

type Add_video_resquest struct {
	Play_url    string `form:"play_url"`
	Cover_url   string `form:"cover_url"`
	Description string `form:"description"`
	Title       string `form:"title"`
	Line_type   int    `form:"line_type"`
	MiniType    int    `form:"mini_type"`
	StageType   int    `form:"stage_type"`
}

// 跟删除操作响应共用
type Add_video_response struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type Delete_video_resquest struct {
	Vid uint `form:"vid"`
}
