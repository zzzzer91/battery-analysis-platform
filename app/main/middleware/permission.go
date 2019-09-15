package middleware

import (
	"battery-analysis-platform/app/main/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PermissionRequired(permission int) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userName := session.Get("userName")
		if userName == nil {
			c.AbortWithStatus(401)
			return
		}

		user, err := model.GetUser(userName.(string))
		if err != nil {
			c.AbortWithError(401, err)
			return
		}
		if user.Type < permission {
			c.AbortWithStatus(403)
			return
		}

		c.Next()
	}
}
