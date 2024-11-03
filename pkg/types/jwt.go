package types

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	jwt.StandardClaims
	UserID    string `json:"user_id"`
	UserEmail string `json:"user_email"`
}
