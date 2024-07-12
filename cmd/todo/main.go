package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"task2/internal/config"
	"task2/internal/router"
)

func main() {
	cfg := config.MustLoadConfig()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	log.Println(cfg)

	logger.Info(cfg.HttpServer.Port)
	logger.Info(cfg.HttpServer.Host)

	r := router.Router()

	err := http.ListenAndServe(":"+cfg.HttpServer.Port, r)
	if err != nil {
		logger.Error("Error starting server: %v", err)
	}
}
