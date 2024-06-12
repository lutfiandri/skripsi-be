package rest

import "github.com/golang-jwt/jwt"

type JWTUserClaimsData struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type JWTClaims struct {
	jwt.StandardClaims
	User        JWTUserClaimsData `json:"user"`
	ClientId    *string           `json:"client_id"`
	Permissions []string          `json:"permissions"`
}

type JWTRefreshClaims struct {
	jwt.StandardClaims
	User     JWTUserClaimsData `json:"user"`
	ClientId *string           `json:"client_id"`
}
