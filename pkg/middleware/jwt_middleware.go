package middleware

import (
	"car-rental-backend/pkg/jwt"
	"car-rental-backend/pkg/response"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
	UserIDKey           = "user_id"
	UserEmailKey        = "user_email"
)

func JWTMiddleware(jwtManager *jwt.JWTManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get(AuthorizationHeader)
		if authHeader == "" {
			return response.Failed(c, fiber.StatusUnauthorized, "Missing authorization header", nil)
		}

		if !strings.HasPrefix(authHeader, BearerPrefix) {
			return response.Failed(c, fiber.StatusUnauthorized, "Invalid authorization format. Use: Bearer <token>", nil)
		}

		tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
		if tokenString == "" {
			return response.Failed(c, fiber.StatusUnauthorized, "Missing token", nil)
		}

		claims, err := jwtManager.ValidateToken(tokenString)
		if err != nil {
			return response.Failed(c, fiber.StatusUnauthorized, "Invalid or expired token", nil)
		}

		c.Locals(UserIDKey, claims.UserID)
		c.Locals(UserEmailKey, claims.Email)

		return c.Next()
	}
}

func GetUserID(c *fiber.Ctx) (uuid.UUID, bool) {
	userID, ok := c.Locals(UserIDKey).(uuid.UUID)
	return userID, ok
}

func GetUserEmail(c *fiber.Ctx) (string, bool) {
	email, ok := c.Locals(UserEmailKey).(string)
	return email, ok
}
