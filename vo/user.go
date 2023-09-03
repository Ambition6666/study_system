package vo

type UserInfo struct {
	ID               int64  `json:"uid" gorm:"primarykey"`
	NickName         string `json:"nickname"`
	Email            string `json:"email"`
	Avatar           string `json:"avatar"`
	IndividualResume string `json:"individual_resume"`
}
type Get_user_info_response struct {
	Code int
	UserInfo
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
