package routers

import (
	"api-hateoas/handlers"
	"api-hateoas/middleware"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	
	
	// public routes
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	
	// protected routes
	protected := router.PathPrefix("/protected").Subrouter()
	protected.HandleFunc("", handlers.ProtectedHandler).Methods("GET")
	protected.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	protected.Use(middleware.AuthMiddleware)

	return router
}