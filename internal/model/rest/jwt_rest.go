package rest

import (
	"github.com/golang-jwt/jwt"
)

type JWTUserClaimsData struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type JWTClaims struct {
	jwt.StandardClaims
	User JWTUserClaimsData `json:"user"`
}

type JWTRefreshClaims struct {
	jwt.StandardClaims
	User JWTUserClaimsData `json:"user"`
}
