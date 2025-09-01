package Auth

import "github.com/gin-gonic/gin"

type Auth struct {
}

func (u *Auth) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{
			"username": "admin",
			"token":    "123456",
		},
	})
}
