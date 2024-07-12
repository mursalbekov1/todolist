package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"task2/internal/handlers"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Get("/getTask", handlers.GetTask)
		r.Get("/updateTask", handlers.GetTasks)
		r.Post("/updateTask", handlers.UpdateTask)
		r.Post("/addTask", handlers.AddTask)
		r.Delete("/getTasks", handlers.DeleteTask)

	})

	return router
}
