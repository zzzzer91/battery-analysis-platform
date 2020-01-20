package middleware

import (
	"battery-analysis-platform/app/main/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PermissionRequired(permission int) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userName := session.Get("userName")
		if userName == nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		user, err := model.GetUserFromCache(userName.(string))
		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}
		// 保存用户信息到本次会话中
		c.Set("user", user)
		if user.Type < permission {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
