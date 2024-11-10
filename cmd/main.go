package main

import (
	"backend-vtb/internal/config"
	"backend-vtb/internal/http"
	"backend-vtb/internal/repository"
	"backend-vtb/internal/server"
	"backend-vtb/internal/service"
	"backend-vtb/pkg/auth"
	"backend-vtb/pkg/database"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title LinkBase
// @version 1.0
// @description LinkBase API

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey UsersAuth
// @in header
// @name Authorization
func main() {
	cfg := config.MustLoad()

	fmt.Println("Config: ", cfg)

	logger := setupLogger()

	postgresClient, err := database.NewPostgresClient(cfg.Postgres)
	if err != nil {
		log.Fatalf("Failed to initialize Postgres DB: %v", err)
	}

	repos := repository.NewRepository(postgresClient)

	tokenManager, err := auth.NewManager(cfg.JWT.SigningKey)
	if err != nil {
		log.Fatalf("Failed to initialize token manager: %v", err)
	}

	serv := service.NewService(repos, logger)

	handlers := http.NewHandler(serv.Base, tokenManager)

	srv := server.NewServer(cfg.HTTP, handlers.Init())
	go func() {
		if err := srv.Run(); err != nil {
			logger.Info("shutting down server", slog.String("reason", err.Error()))
		}
	}()

	logger.Info("server started", slog.String("address", cfg.HTTP.Port))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Error("failed to stop server", slog.String("reason", err.Error()))
	}
}

// setupLogger initializes and returns a new logger instance configured
// with a text handler that outputs to the standard output.
// The logger is set to debug level and includes the source of the log.
func setupLogger() *slog.Logger {
	var logger *slog.Logger

	// Create a text handler for logging, outputting to standard output
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	logger = slog.New(handler)

	return logger
}
