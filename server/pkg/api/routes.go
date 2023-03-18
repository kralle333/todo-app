package api

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	todoIDKey = "todoID"
	taskIDKey = "taskID"
)

func (a *todoApp) routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Origin",
			"Accept",
			"X-Requested-With",
			"Content-Type",
			"Access-Control-Request-Method",
			"Access-Control-Request-Headers"},
		//	ExposedHeaders:   []string{"Link"},
		//	AllowCredentials: false,
		MaxAge: 300, // Maximum value not ignored by any of major browsers
	}))

	r.HandleFunc("/v1/healthcheck", a.healthcheckHandler)
	r.Route("/v1/todos", func(r chi.Router) {
		r.Get("/", a.getAllTodos)
		r.Post("/", a.createTodoList)
		r.Route(fmt.Sprintf("/{%s}", todoIDKey), func(r chi.Router) {
			r.Get("/", a.getTodo)
			r.Delete("/", a.deleteTodo)
			r.Get("/tasks", a.getAllTasks)
			r.Post("/tasks", a.addTask)
		})
	})

	r.Route(fmt.Sprintf("/v1/tasks/{%s}", taskIDKey), func(r chi.Router) {
		r.Post("/complete", a.markCompleted)
		r.Delete("/", a.deleteTask)
	})
	return r
}
