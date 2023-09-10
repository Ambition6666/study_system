package vo

import "studysystem/models"

type Add_video_resquest struct {
	Play_url    string  `form:"play_url"`
	Cover_url   string  `form:"cover_url"`
	Description string  `form:"description"`
	Title       string  `form:"title"`
	Type        float64 `form:"type"`
}

//跟删除响应共用
type Add_video_response struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type Delete_video_resquest struct {
	Vid uint `form:"vid"`
}

type Video_list_resquest struct {
	Type      float64 `form:"type"`
	Line_type int     `form:"line_type"`
	Offset    int64   `form:"offset"`
}
type Video_list_response struct {
	Code        int                `json:"code"`
	Videos      []models.Video     `json:"videos"`
	Study_route models.Study_route `json:"study_route"`
	Type        float64            `json:"type"`
}
