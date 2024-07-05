package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"task2/internal/config"
)

func main() {
	cfg := config.MustLoadConfig()

	log.Println(cfg)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/healthCheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(cfg.HttpServer.Host+":"+cfg.HttpServer.Port, r)

}
