package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sreio/gold/config"
	webrouter "github.com/sreio/gold/web/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var R *gin.Engine

func Start(cfg config.Web) error {
	log.Println("启动 Web Api 服务...")
	gin.SetMode(gin.ReleaseMode)
	R = gin.Default()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: webrouter.NewRouter(R).Handler(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("服务启动成功：http://%s:%d\n", cfg.Host, cfg.Port)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no params) by default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("关闭服务 ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("服务关闭失败:", err)
	}
	log.Println("服务已关闭")

	return nil
}

func StartCron() {

}
