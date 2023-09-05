package vo

import "studysystem/models"

type Get_user_info_response struct {
	Code      int         `json:"code"`
	User_info models.User `json:"user_info"`
}
type Update_user_request struct {
	Data   string `form:"data"`
	Action int    `form:"action"`
}
type Update_user_response struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}
type Delete_user_response struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}
