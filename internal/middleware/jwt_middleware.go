package middleware

import (
	"slices"
	"strings"

	"skripsi-be/internal/interface/rest"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
)

const (
	CtxClaims = "claims"
)

func NewAuthenticator() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		authHeaderList := strings.Split(authHeader, " ")

		if len(authHeaderList) < 2 || authHeaderList[0] != "Bearer" || authHeaderList[1] == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or missing Bearer token")
		}

		tokenString := authHeaderList[1]

		claims, err := helper.ParseJwt(tokenString)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		c.Locals(CtxClaims, claims)
		return c.Next()
	}
}

func NewAuthorizer(permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals(CtxClaims).(rest.JWTClaims)

		if !slices.Contains(claims.Permissions, permission) {
			return fiber.NewError(fiber.StatusForbidden, "Doesn't have permission to access this resource")
		}

		return c.Next()
	}
}
