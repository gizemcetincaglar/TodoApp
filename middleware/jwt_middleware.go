package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

var jwtSecret = []byte("gizli-anahtar") 

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header missing"})
		}

		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("userId", claims["id"])
		c.Locals("role", claims["role"])
		c.Locals("username", claims["username"])

		return c.Next()
	}
}
