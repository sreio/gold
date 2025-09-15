package Auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sreio/gold/web/service/jwt"
	"net/http"
	"strings"
)

type Auth struct {
}

func (u *Auth) Login(c *gin.Context) {

	apiToken, ok := c.GetPostForm("api_token")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": "401",
			"msg":  "request api_token is null",
		})
		return
	}

	if len(apiToken) < 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": "401",
			"msg":  "request api_token is null",
		})
		return
	}

	CfgToken, ok := c.Get("api_token")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": "401",
			"msg":  "config api_token is null",
		})
		return
	}

	if !strings.EqualFold(CfgToken.(string), apiToken) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": "401",
			"msg":  "api_token is error",
		})
		return
	}

	token, err := jwt.GenToken(apiToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "gen token is error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{
			"token": token,
		},
	})
}
