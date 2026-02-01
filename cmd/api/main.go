package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Vinaychinnu/platform-api/internal/infrastructure"
	apphttp "github.com/Vinaychinnu/platform-api/internal/transport/http"
	"github.com/Vinaychinnu/platform-api/pkg/config"
)

func main() {
	cfg := config.Load()

	db, err := infrastructure.OpenPostgres(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
	)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: apphttp.NewRouter(db),
	}

	go func() {
		log.Printf("starting server on :%s", cfg.HTTPPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown error: %v", err)
	}

	log.Println("server stopped")
}
