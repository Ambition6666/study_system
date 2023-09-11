package vo

type Add_study_route_resquest struct {
	Line_type   int    `json:"line_type"`
	Description string `json:"description"`
}
type Add_study_route_response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
