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
	// 会议列表
	auth.GET("/meeting/list", service.MeetingList)
	// 创建会议
	auth.POST("/meeting/create", service.MeetingCreate)
	// 编辑会议
	auth.PUT("/meeting/edit", service.MeetingEdit)
	// 删除会议
	auth.DELETE("/meeting/delete", service.MeetingDelete)

	return r
}
