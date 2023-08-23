package apis

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spearexit/dropick.core/v2/internal/shared"
	"github.com/spearexit/dropick.core/v2/api"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

type App struct {
	Db       *gorm.DB
	Validate *validator.Validate
}

//	@BasePath	/api/v1

func New() *App {
	db, err := shared.SetConnection(shared.Config.Database)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection established")
	}

	return &App{
		Db:       db,
		Validate: validator.New(),
	}
}

func (a *App) Close() {
	if sqlDB, err := a.Db.DB(); err == nil {
		sqlDB.Close()
	}
}

func (a *App) GetRouter() *gin.Engine {
	var r *gin.Engine
	if shared.Config.Server.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		r = gin.Default()
	}
	api.SwaggerInfo.BasePath = "/api/v1"

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
		Addr:    ":" + shared.Config.Server.Port,
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
