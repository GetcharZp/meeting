package service

import (
	"github.com/gin-gonic/gin"
	"meeting/internal/helper"
	"meeting/internal/models"
	"net/http"
	"time"
)

func MeetingList(c *gin.Context) {
	in := new(MeetingListRequest)
	err := c.ShouldBindQuery(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	var list []*MeetingListReply
	var cnt int64
	tx := models.DB.Model(&models.RoomBasic{})
	if in.Keyword != "" {
		tx.Where("name LIKE ?", "%"+in.Keyword+"%")
	}
	err = tx.Count(&cnt).Limit(in.Size).Offset((in.Page - 1) * in.Size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常：" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":  list,
			"count": cnt,
		},
	})
}

func MeetingCreate(c *gin.Context) {
	uc := c.MustGet("user_claims").(*helper.UserClaims)
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
		BeginAt:  time.UnixMilli(in.BeginAt),
		EndAt:    time.UnixMilli(in.EndAt),
		CreateId: uc.Id,
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

func MeetingEdit(c *gin.Context) {
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	in := new(MeetingEditRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	err = models.DB.Model(new(models.RoomBasic)).Where("identity = ? AND create_id = ?", in.Identity, uc.Id).
		Updates(map[string]any{
			"name":     in.Name,
			"begin_at": time.UnixMilli(in.BeginAt),
			"end_at":   time.UnixMilli(in.EndAt),
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

func MeetingDelete(c *gin.Context) {
	identity := c.Query("identity")
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	err := models.DB.Where("identity = ? AND create_id = ?", identity, uc.Id).Delete(&models.RoomBasic{}).Error
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
