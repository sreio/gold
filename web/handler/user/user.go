package user

import "github.com/gin-gonic/gin"

type User struct {
}

func (u *User) List(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{},
	})
}

func (u *User) Add(c *gin.Context) {
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
