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
		var userLoginService service.UserLoginService
		// ShouldBind() 会检测是否满足设置的 bind 标签要求
		if err := c.ShouldBindJSON(&userLoginService); err != nil {
			c.AbortWithStatus(500)
			return
		} else {
			if user, err := userLoginService.Login(); err != nil {
				code = jd.ERROR
				msg = err.Error()
			} else {
				// 设置Session
				s := sessions.Default(c)
				s.Clear()
				s.Set("user_id", user.ID)
				_ = s.Save()
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
	s := sessions.Default(c)
	s.Clear()
	_ = s.Save()
	res := jd.Build(jd.SUCCESS, "", nil)
	c.JSON(200, res)
}
