package router

import (
	"github.com/gin-gonic/gin"
	"meeting/internal/server/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// meeting
	r.POST("/meeting/create", service.MeetingCreate)

	// 用户登录
	r.POST("/user/login", service.UserLogin)

	return r
}
