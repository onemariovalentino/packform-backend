package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"packform-backend/src/app/orders/handlers"
	"packform-backend/src/pkg/config"
	"packform-backend/src/pkg/di"
	"syscall"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

const DefaultPort = 8080

type (
	Server struct {
		handler *gin.Engine
	}
)

func New() *Server {
	switch config.Env.Environment {
	default:
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "production":
		gin.SetMode(gin.ReleaseMode)
	}

	services := di.NewDependency()

	g := gin.Default()
	g.Use(gin.Recovery())
	g.Use(requestid.New())
	orderHTTPHandler := handlers.NewOrderHTTPHandler(services.OrderUsecase)
	orderHTTPHandler.Mount(g)

	return &Server{handler: g}
}

func (s *Server) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	port := config.Env.Port
	if port == 0 {
		port = DefaultPort
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
