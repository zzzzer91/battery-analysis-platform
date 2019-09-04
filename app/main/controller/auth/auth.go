package auth

import (
	"battery-anlysis-platform/app/main/service"
	"battery-anlysis-platform/pkg/jd"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var code int
	var msg string
	var data interface{}

	if c.Request.Method == "GET" {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId != nil {
			if user, err := service.LoginByCookie(userId.(int)); err != nil {
				code = jd.ERROR
				msg = err.Error()
			} else {
				code = jd.SUCCESS
				data = user
			}
		}
	} else if c.Request.Method == "POST" {
		var s service.UserLoginService
		// ShouldBind() 会检测是否满足设置的 bind 标签要求
		if err := c.ShouldBindJSON(&s); err != nil {
			c.AbortWithError(500, err)
			return
		} else {
			if user, err := s.Login(); err != nil {
				code = jd.ERROR
				msg = err.Error()
			} else {
				// 设置Session
				session := sessions.Default(c)
				session.Clear()
				session.Set("user_id", user.ID)
				_ = session.Save()
				//
				code = jd.SUCCESS
				data = user
			}
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
