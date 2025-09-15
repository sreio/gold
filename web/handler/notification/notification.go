package notification

import (
	"github.com/gin-gonic/gin"
	"github.com/sreio/gold/web/model"
)

type Notification struct {
}

func (s *Notification) List(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{
			"list": model.ChanelDataList,
		},
	})
}
