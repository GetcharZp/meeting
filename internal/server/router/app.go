package router

import (
	"github.com/gin-gonic/gin"
	"meeting/internal/middlewares"
	"meeting/internal/server/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(middlewares.Cors())
	// 用户登录
	r.POST("/user/login", service.UserLogin)

	auth := r.Group("/auth", middlewares.Auth())

	// meeting
	// 创建会议
	auth.POST("/meeting/create", service.MeetingCreate)

	return r
}
