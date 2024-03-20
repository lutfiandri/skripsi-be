package middleware

import (
	"strings"

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
