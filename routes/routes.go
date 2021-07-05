package routes

import (
	"github.com/gorilla/mux"
	"github.com/ec965/todo-api/handlers"
)

func Init(r *mux.Router){
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	api.HandleFunc("/ping", handlers.Ping).Methods("GET")
}