package middleware

import (
	"battery-analysis-platform/app/main/dao"
	"battery-analysis-platform/app/main/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PermissionRequired(permission int) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get("userId")
		if userId == nil {
			c.AbortWithStatus(401)
			return
		}
		var user model.User
		err := dao.MysqlDB.First(&user, userId).Error
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		if user.Type < permission {
			c.AbortWithStatus(403)
			return
		}

		c.Next()
	}
}
