package routes

import (
	"github.com/gorilla/mux"

	"github.com/ec965/todo-api/handlers"
)

func Init(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	private := api.PathPrefix("/private").Subrouter()
	private.Use(handlers.JwtMiddleWare)
	private.HandleFunc("/ping", handlers.Ping).Methods("GET")

	task := api.PathPrefix("/task").Subrouter()
	task.Use(handlers.JwtMiddleWare)
	task.HandleFunc("", handlers.GetTasks).Methods("GET")
	task.HandleFunc("", handlers.CreateTask).Methods("POST")

	// auth does not require a token
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", handlers.Signup).Methods("POST")
	auth.HandleFunc("/login", handlers.Login).Methods("POST")
	auth.HandleFunc("/ping", handlers.Ping).Methods("GET")
}
