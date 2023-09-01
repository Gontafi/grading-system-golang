package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

const (
	TeacherRoleId = 1
	StudentRoleId = 2
	AdminRoleId   = 3
	GuestRoleId   = 4
)

func (h *Handler) AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		headerParts := strings.Split(tokenString, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		_, err := ParseToken(headerParts[1])
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		return c.Next()
	}
}

func (h *Handler) TeacherRoleMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		headerParts := strings.Split(tokenString, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		claims, err := ParseToken(headerParts[1])
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}
		if claims.RoleID == AdminRoleId {
			return c.Next()
		}
		if claims.RoleID != TeacherRoleId {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}

		return c.Next()
	}
}

func (h *Handler) AdminRoleMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		headerParts := strings.Split(tokenString, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		claims, err := ParseToken(headerParts[1])
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		if claims.RoleID != AdminRoleId {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}

		return c.Next()
	}
}
