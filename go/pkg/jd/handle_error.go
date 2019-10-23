// 封装错误处理

package jd

func HandleError(err error) (code int, msg string) {
	if err != nil {
		code = ERROR
		msg = err.Error()
	} else {
		code = SUCCESS
	}
	return
}
