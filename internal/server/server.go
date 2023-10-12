package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/automagic-tools/go-coding-challenge/SeeCV/config"
	"github.com/automagic-tools/go-coding-challenge/SeeCV/internal/db"

	"github.com/automagic-tools/go-coding-challenge/SeeCV/internal/apis"
	"github.com/automagic-tools/go-coding-challenge/SeeCV/utils/logger"
	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
}

func (s *Server) Run() error {
	logger.Info("Starting the server on", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("Server stopped with error", err)
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	logger.Info("Shutting down server...")
	return s.server.Shutdown(ctx)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/upload-resume", apis.UploadResumeHandler)
	router.DELETE("/candidate/:id", apis.DeleteCandidateHandler)
	router.GET("/candidate/:id", apis.GetCandidateHandler)
	router.GET("/candidate/:id/skills", apis.GetSkillsForCandidateHandler)
	router.GET("/candidate/:email", apis.GetCandidateByEmailHandler)

	return router
}

func NewServer(port string, router *gin.Engine) *Server {
	return &Server{
		server: &http.Server{
			Addr:    port,
			Handler: router,
		},
	}
}

func StartServer() {

	if err := config.LoadConfig(); err != nil {
		logger.Error("Failed to load configuration", err)
	}

	if err := db.StartDB(); err != nil {
		logger.Error("Error connecting to database", err)
		return
	}
	logger.Info("Connected to the database")

	router := setupRouter()
	s := NewServer(config.Vauban.Server.Port, router)
	go func() {
		if err := s.Run(); err != nil {
			logger.Error("Server stopped with error", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Blocking until termination signal is received
	<-quit

	logger.Info("Shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", err)
	}

	if err := db.CloseDB(); err != nil {
		logger.Error("Failed to close the database", err)
	}
	logger.Info("Server exited properly")

}
