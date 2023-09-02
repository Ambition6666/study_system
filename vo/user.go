package vo

type UserInfo struct {
	ID               int64  `json:"uid" gorm:"primarykey"`
	NickName         string `json:"nickname"`
	Email            string `json:"email"`
	Avatar           string `json:"avatar"`
	IndividualResume string `json:"individual_resume"`
}
type Update_user_request struct {
	Data   string `form:"data"`
	Action int    `form:"action"`
}
