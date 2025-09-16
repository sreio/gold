package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sreio/gold/web/dto"
	"github.com/sreio/gold/web/repository"
	"github.com/sreio/gold/web/service/data"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type User struct {
	svc *data.UserService
}

func NewUser(db *gorm.DB) *User {
	return &User{
		svc: data.NewUserService(repository.NewUserRepo(db)),
	}
}

func (u *User) List(c *gin.Context) {
	var q dto.QueryUser
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "1", "msg": err.Error()})
		return
	}
	list, total, err := u.svc.List(q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "1", "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "0",
		"msg":  "success",
		"data": gin.H{
			"list":  list,
			"total": total,
			"page":  q.Page,
			"size":  q.Size,
		},
	})
}

func (u *User) Add(c *gin.Context) {
	var userDto dto.CreateUserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "1", "msg": "参数错误: " + err.Error()})
		return
	}

	exits, err := u.svc.Exits(userDto.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "1", "msg": "查询错误: " + err.Error()})
		return
	}

	if exits {
		c.JSON(http.StatusBadRequest, gin.H{"code": "1", "msg": "用户已存在"})
		return
	}

	created, err := u.svc.Create(userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "1", "msg": "添加失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "0", "msg": "success", "data": created})
}

func (u *User) Edit(c *gin.Context) {
	id64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": "1", "msg": "id 无效"})
		return
	}
	var userDto dto.UpdateUserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "1", "msg": "参数错误: " + err.Error()})
		return
	}
	if err := u.svc.Update(uint(id64), userDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "1", "msg": "更新失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "0", "msg": "success"})
}

func (u *User) Del(c *gin.Context) {
	id64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": "1", "msg": "id 无效"})
		return
	}
	if err := u.svc.Delete(uint(id64)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "1", "msg": "删除失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "0", "msg": "success"})
}
