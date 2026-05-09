package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"urlshortener/internal/config"
	"urlshortener/internal/http-serv/handlers/redirect"
	"urlshortener/internal/http-serv/handlers/url/save"
	"urlshortener/internal/lib/logger/sl"
	"urlshortener/internal/storage/postgres"
)

const (
	envLocal = "local"
)

func main() {
	cfg := config.Load()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)
	log.Info("bebebe", slog.String("env", cfg.Env))

	storage, err := postgres.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to initialize storage", sl.Err(err))
		return
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	workDir, err := os.Getwd()
	if err != nil {
		log.Error("failed to get working directory", sl.Err(err))
		return
	}

	staticDir := filepath.Join(workDir, "static")
	if err := os.MkdirAll(staticDir, 0755); err != nil {
		log.Error("failed to create static directory", sl.Err(err))
		return
	}

	fileServer := http.FileServer(http.Dir(staticDir))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		indexPath := filepath.Join(workDir, "index.html")
		if _, err := os.Stat(indexPath); os.IsNotExist(err) {
			http.Error(w, "index.html not found", http.StatusNotFound)
			return
		}
		http.ServeFile(w, r, indexPath)
	})

	router.Post("/api/url", save.New(log, storage))

	router.Get("/{alias}", redirect.New(log, storage))

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error("failed to start server", sl.Err(err))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return log
}
