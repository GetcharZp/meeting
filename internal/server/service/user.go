package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"meeting/internal/helper"
	"meeting/internal/models"
	"net/http"
)

func UserLogin(c *gin.Context) {
	in := new(UserLoginRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	if in.Username == "" || in.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填信息为空",
		})
		return
	}
	in.Password = helper.GetMd5(in.Password)

	data := new(models.UserBasic)
	err = models.DB.Where("username = ? AND password = ? ", in.Username, in.Password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get UserBasic Error:" + err.Error(),
		})
		return
	}

	token, err := helper.GenerateToken(data.ID, data.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GenerateToken Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
