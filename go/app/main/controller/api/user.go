package api

import (
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var s service.UserCreateService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	res, err := s.Do()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, res)
}

func ModifyUser(c *gin.Context) {
	var s service.UserModifyService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	s.UserName = c.Param("name")
	res, err := s.Do()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, res)
}

func ListUser(c *gin.Context) {
	var s service.UserListService
	res, err := s.Do()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, res)
}
