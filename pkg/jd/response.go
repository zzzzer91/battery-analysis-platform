package jd

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Build(code int, msg string, data interface{}) *response {
	return &response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
