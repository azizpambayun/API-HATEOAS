package handlers

import (
	"api-hateoas/config"
	"api-hateoas/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// In-memory user store
var users = map[string]string{}

// Claims represents JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Response with HATEOAS links
type Response struct {
	Data interface{} `json:"data"`
	Links map[string]string `json:"links"`
}

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.User
    body, err := io.ReadAll(r.Body) // Read raw body for debugging
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
        return
    }
    log.Printf("Received body: %s\n", body) // Log raw body for debugging

    err = json.Unmarshal(body, &creds) // Unmarshal into creds
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    log.Printf("Parsed creds: %+v\n", creds) // Log the parsed creds

    if creds.Username == "" || creds.Password == "" {
        http.Error(w, "Username and password are required", http.StatusBadRequest)
        return
    }


	if creds.Username == "" || creds.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	if _, exists := users[creds.Username]; exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	users[creds.Username] = string(hashedPassword)

	response := Response{
		Data: map[string]string{
			"message": "User registered successfully",
		},
		Links: map[string]string{
			"login": "/login",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.User

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	storedPassword, exists := users[creds.Username]
	if !exists {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Create JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.Key())
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: map[string]string{
			"token": tokenString,
		},
		Links: map[string]string{
			"protected": "/protected",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}