package vo

import "studysystem/models"

type Add_video_resquest struct {
	Play_url    string `form:"play_url"`
	Cover_url   string `form:"cover_url"`
	Description string `form:"description"`
	Title       string `form:"title"`
	Type        int    `form:"type"`
}
type Add_video_response struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
type Video_list_response struct {
	Videos    []models.Video `json:"videos"`
	Type      int            `json:"type"`
	Line_type int            `json:"line_type"`
}
