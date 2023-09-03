package vo

type Register_resquest struct {
	Email string `form:"email"`
	Pwd   string `form:"pwd"`
}

type Register_response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type Login_resquest struct {
	Email string `form:"email"`
	Pwd   string `form:"pwd"`
}
type Login_response struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
}
