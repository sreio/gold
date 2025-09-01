package task

import "github.com/gin-gonic/gin"

type Task struct {
}

func (u *Task) List(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{},
	})
}

func (u *Task) Add(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{},
	})
}

func (u *Task) Edit(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{},
	})
}

func (u *Task) Del(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{},
	})
}
