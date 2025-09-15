package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sreio/gold/database"
	"github.com/sreio/gold/web/model"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
}

func (u *User) List(c *gin.Context) {
	var ApiRequestListUser []model.ApiRequestListUser
	database.DB.Preload("UserConf").Find(&ApiRequestListUser)
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": ApiRequestListUser,
	})
}

func (u *User) Add(c *gin.Context) {
	var ApiRequestAddUser model.ApiRequestAddUser
	if err := c.ShouldBindJSON(&ApiRequestAddUser); err != nil {
		c.JSON(200, gin.H{
			"code": "1",
			"msg":  "参数错误",
			"data": gin.H{},
		})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		userModel := &model.User{
			Name:    ApiRequestAddUser.Name,
			Cron:    ApiRequestAddUser.Cron,
			SaveDay: ApiRequestAddUser.SaveDay,
		}
		err := tx.Create(userModel).Error
		if err != nil {
			return err
		}
		for _, v := range ApiRequestAddUser.UserConf {
			userConfModel := &model.UserConf{
				UserID: userModel.ID,
				Key:    v.Key,
				Type:   v.Type,
				Value:  v.Value,
			}
			err := tx.Create(userConfModel).Error
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": "1",
			"msg":  "添加失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{},
	})
}

func (u *User) Edit(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{},
	})
}

func (u *User) Del(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{},
	})
}
