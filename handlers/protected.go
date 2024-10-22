package handlers

import (
	"api-hateoas/middleware"
	"encoding/json"
	"fmt"
	"net/http"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(middleware.UsernameKey).(string)

	response := Response{
		Data: map[string]string{
			"message": fmt.Sprintf("Hello, %s! this is protected resources", username),
		},
		Links: map[string]string{
			"logout": "/protected/logout",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Data: map[string]string{
			"message": "Logged out successfully",
		},
		Links: map[string]string{
			"login": "/login",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}