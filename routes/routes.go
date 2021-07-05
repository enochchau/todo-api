package routes

import (
	"github.com/gorilla/mux"

	"github.com/ec965/todo-api/handlers"
	"github.com/ec965/todo-api/handlers/middleware"
)

func Init(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	private := api.PathPrefix("/private").Subrouter()
	private.Use(middleware.Jwt)
	private.HandleFunc("/ping", handlers.Ping).Methods("GET")

	// auth does not require a token
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	auth.HandleFunc("/login", handlers.Login).Methods("POST")
	auth.HandleFunc("/ping", handlers.Ping).Methods("GET")
}
