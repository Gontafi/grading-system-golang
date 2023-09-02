package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
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

		_, err := h.Service.ParseToken(headerParts[1])
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

		claims, err := h.Service.ParseToken(headerParts[1])
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		role, err := h.Service.GetUserRole(claims.UserID)
		if err != nil {
			return err
		}
		if role.ID == AdminRoleId {
			return c.Next()
		}
		if role.ID != TeacherRoleId {
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

		claims, err := h.Service.ParseToken(headerParts[1])
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		role, err := h.Service.GetUserRole(claims.UserID)
		if err != nil {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Failed to get role"})
		}

		if role.ID != AdminRoleId {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}

		return c.Next()
	}
}

func (h *Handler) MarkBelongsTeacherMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		markIDParam := c.Params("id")
		markID, err := strconv.Atoi(markIDParam)
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid mark ID"})
		}

		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		headerParts := strings.Split(tokenString, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		claims, err := h.Service.ParseToken(headerParts[1])
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		existingMark, err := h.Service.GetMarkByID(markID)
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve mark"})
		}

		if claims.UserID != existingMark.TeacherID {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}

		return c.Next()
	}
}
