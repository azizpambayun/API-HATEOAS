package models

import "github.com/golang-jwt/jwt"

// User represents a user in the system
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}