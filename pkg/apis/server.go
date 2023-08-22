package apis

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"github.com/spearexit/dropick.core/v2/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

type App struct {
	Db       *gorm.DB
	Validate *validator.Validate
}

type JSONResult struct {
    Code    int          `json:"code" `
    Message string       `json:"message"`
    Data    interface{}  `json:"data"`
}

// @BasePath /api/v1

// Ping godoc
// @Summary ping to check server status
// @Schemes
// @Description do ping
// @Tags /health
// @Accept json
// @Produce json
// @Success 200 {object} JSONResult 
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, JSONResult{
		Code:    200,
		Message: "pong",
		Data:    nil,
	})
}

func (a *App) Close() {
	if sqlDB, err := a.Db.DB(); err == nil {
		sqlDB.Close()
	}
}

func (a *App) GetRouter() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		{
			health.GET("/ping", Ping)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

func (a *App) Start() {
	router := a.GetRouter()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown Server ...")

	// close db and socket connection
	a.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	log.Println("timeout of 5 seconds.")
	<-ctx.Done()

	log.Println("Server exiting")
}
