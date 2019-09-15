package jd

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Build(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Err(msg string) *Response {
	return &Response{
		Code: ERROR,
		Msg:  msg,
		Data: nil,
	}
}
