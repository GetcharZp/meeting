package service

import (
	"github.com/gin-gonic/gin"
	"meeting/internal/helper"
	"meeting/internal/models"
	"net/http"
	"time"
)

func MeetingCreate(c *gin.Context) {
	in := new(MeetingCreateRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	err = models.DB.Create(&models.RoomBasic{
		Identity: helper.GetUUID(),
		Name:     in.Name,
		BeginAt:  time.UnixMilli(in.CreateAt),
		EndAt:    time.UnixMilli(in.EndAt),
		CreateId: 0, // todo: 从 auth 中间件中获取用户的信息
	}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常：" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}