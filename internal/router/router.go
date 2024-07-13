package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"task2/internal/handlers"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Get("/getTask", handlers.GetTask)
		r.Get("/getTasks", handlers.GetTasks)
		r.Put("/updateTask", handlers.UpdateTask)
		r.Post("/addTask", handlers.AddTask)
		r.Delete("/deleteTask", handlers.DeleteTask)
		r.Get("/healthCheck", healthCheckHandler)
	})

	return router
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
