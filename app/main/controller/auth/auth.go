package auth

import (
	"battery-analysis-platform/app/main/model"
	"battery-analysis-platform/app/main/service"
	"battery-analysis-platform/pkg/jd"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	code := jd.ERROR
	msg := ""
	var data *model.User
	var err error

	if c.Request.Method == "GET" {
		session := sessions.Default(c)
		userId := session.Get("userId")
		if userId != nil {
			data, err = service.LoginByCookie(userId.(int))
			code, msg = jd.HandleError(err)
		}
	} else if c.Request.Method == "POST" {
		var s service.UserLoginService
		// ShouldBind() 会检测是否满足设置的 bind 标签要求
		if err := c.ShouldBindJSON(&s); err != nil {
			c.AbortWithError(500, err)
			return
		}

		data, err = s.Login()
		code, msg = jd.HandleError(err)
		if code == jd.SUCCESS {
			// 设置Session
			session := sessions.Default(c)
			session.Clear()
			session.Set("userId", data.Id)
			_ = session.Save()
		}
	}

	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()
	res := jd.Build(jd.SUCCESS, "", nil)
	c.JSON(200, res)
}
