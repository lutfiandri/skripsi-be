package helper

import (
	"time"

	"skripsi-be/internal/config"
	"skripsi-be/internal/model/rest"

	"github.com/golang-jwt/jwt"
)

func GenerateJwt(userClaimsData rest.JWTUserClaimsData) (string, error) {
	claims := rest.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		User: userClaimsData,
	}

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokens.SignedString([]byte(config.JWTSecretKey))
}

func ParseJwt(tokenString string) (rest.JWTClaims, error) {
	var claims rest.JWTClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecretKey), nil
	})
	if err != nil || !token.Valid {
		return claims, err
	}

	return claims, err
}
