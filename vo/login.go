package vo

type Register_resquest struct {
	Email string `form:"email"`
	Pwd   string `form:"pwd"`
	Code  string `form:"authcode"`
}

type Register_response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type Login_resquest struct {
	Email string `form:"email"`
	Pwd   string `form:"pwd"`
}
type Login_by_auth_code_request struct {
	Email    string `form:"email"`
	AuthCode string `form:"authcode"`
}
type Login_response struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
}
