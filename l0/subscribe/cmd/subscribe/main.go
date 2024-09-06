package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sub/internal/api/cash"
	"sub/internal/api/nats"
	"sub/internal/config"
	"sub/internal/http/get"
	"sub/internal/storage/post"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info(
		"starting publisher-service",
		slog.String("env", cfg.Env),
		slog.String("version", "123"),
	)

	log.Debug("debug messages are enabled")
	context := context.Background()
	//
	sc := *nats.New(cfg.ClusterID, cfg.ClientID, cfg.NatsURL)
	defer sc.Close()

	pg, err := post.NewPG(context, cfg.StoragePath)
	if err != nil {
		panic("DB errore")
	}

	cash, err := cash.New(context, pg)
	if err != nil {
		log.Debug("cash not used connect")
	}

	sc.Subscribe("createOrder", nats.MessageHandlerFunc(cash, pg))

	router := chi.NewRouter()
	//d
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/order", get.New(context, log, cash, pg))

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("panic")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	// Если env config неверен, установите настройки prod по умолчанию из соображений безопасности
	default: 
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
