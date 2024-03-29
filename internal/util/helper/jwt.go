package helper

import (
	"time"

	"skripsi-be/internal/config"
	"skripsi-be/internal/domain"
	"skripsi-be/internal/interface/rest"

	"github.com/golang-jwt/jwt"
)

// access token
func GenerateJwt(user domain.User) (string, error) {
	claims := rest.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(), // 1 hour (google's requirement)
		},
		User: rest.JWTUserClaimsData{
			Id:    user.Id.String(),
			Email: user.Email,
			Name:  user.Name,
		},
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

// refresh token
func GenerateRefreshJwt(user domain.User) (string, error) {
	claims := rest.JWTRefreshClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(999999 * time.Hour).Unix(), // doesn't expire (google's requirement)
		},
		User: rest.JWTUserClaimsData{
			Id:    user.Id.String(),
			Email: user.Email,
			Name:  user.Name,
		},
	}
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokens.SignedString([]byte(config.JWTRefreshSecretKey))
}

func ParseRefreshJwt(tokenString string) (rest.JWTRefreshClaims, error) {
	var claims rest.JWTRefreshClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTRefreshSecretKey), nil
	})
	if err != nil || !token.Valid {
		return claims, err
	}

	return claims, err
}
