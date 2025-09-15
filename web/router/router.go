package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sreio/gold/config"
	"github.com/sreio/gold/web/handler"
	jwtService "github.com/sreio/gold/web/service/jwt"
	"net/http"
)

var H *handler.Handler

func init() {
	H = handler.NewHandler()
}

func NewRouter(router *gin.Engine, cfg config.Web) *gin.Engine {
	handleRouter(router, cfg)
	return router
}

func handleRouter(r *gin.Engine, cfg config.Web) {
	api := r.Group("/api")
	api.Use(ConfigMiddleware(cfg))
	api.POST("/auth/login", H.Auth.Login)
	api.Any("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authApi := api.Group("")
	authApi.Use(AuthMiddle)

	authApi.GET("/source/list", H.Source.List)
	authApi.GET("/notification/list", H.Notification.List)

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
	token := c.GetHeader("api_token")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": "401",
			"msg":  "request api_token is null",
		})
	} else {
		var code = 200
		_, err := jwtService.ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = 402
			default:
				code = 403
			}
			c.AbortWithStatusJSON(code, gin.H{
				"code": code,
				"msg":  "api_token is invalid",
			})
		}
	}

	c.Next()
}

func ConfigMiddleware(cfg config.Web) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", cfg)
		c.Set("api_token", cfg.Token)
		c.Next()
	}
}
