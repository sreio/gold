package gold

import "github.com/gin-gonic/gin"

type Gold struct {
}

func (u *Gold) List(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{},
	})
}

func (u *Gold) Add(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{},
	})
}

func (u *Gold) Del(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{},
	})
}
