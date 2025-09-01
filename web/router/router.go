package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sreio/gold/web/handler"
)

var H *handler.Handler

func init() {
	H = handler.NewHandler()
}

func NewRouter(router *gin.Engine) *gin.Engine {
	baseApi(router)
	handleRouter(router)
	return router
}

func baseApi(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func handleRouter(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/auth/login", H.Auth.Login)

	authApi := api.Group("")
	authApi.Use(AuthMiddle)

	authApi.GET("/source/list", H.Source.List)

	userApi := authApi.Group("/user")
	userApi.GET("/list", H.User.List)
	userApi.POST("/add", H.User.Add)
	userApi.POST("/edit", H.User.Edit)
	userApi.POST("/del", H.User.Del)

	goldApi := authApi.Group("/gold")
	goldApi.GET("/list", H.Gold.List)
	goldApi.POST("/add", H.Gold.Add)
	goldApi.POST("/del", H.Gold.Del)

	taskApi := authApi.Group("/task")
	taskApi.GET("/list", H.Task.List)
	taskApi.POST("/add", H.Task.Add)
	taskApi.POST("/edit", H.Task.Edit)
	taskApi.POST("/del", H.Task.Del)
}

func AuthMiddle(c *gin.Context) {

	c.Next()
}
