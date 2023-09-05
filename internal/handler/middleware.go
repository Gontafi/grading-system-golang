package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grading-system-golang/internal/models"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	TeacherRoleId = 1
	StudentRoleId = 2
	AdminRoleId   = 3
)

func (h *Handler) AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")
		if tokenString == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		headerParts := strings.Split(tokenString, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		_, err := h.Service.ParseToken(headerParts[1])
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		return ctx.Next()
	}
}

func (h *Handler) TeacherRoleMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")
		if tokenString == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		headerParts := strings.Split(tokenString, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		claims, err := h.Service.ParseToken(headerParts[1])
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		role, err := h.Service.GetUserRole(claims.UserID)
		if err != nil {
			return err
		}
		if role.ID == AdminRoleId {
			return ctx.Next()
		}
		if role.ID != TeacherRoleId {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}

		return ctx.Next()
	}
}

func (h *Handler) StudentRoleMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")
		if tokenString == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		headerParts := strings.Split(tokenString, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		claims, err := h.Service.ParseToken(headerParts[1])
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		role, err := h.Service.GetUserRole(claims.UserID)
		if err != nil {
			return err
		}
		if role.ID == AdminRoleId {
			return ctx.Next()
		}
		if role.ID != StudentRoleId {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}

		return ctx.Next()
	}
}

func (h *Handler) StudentBelongsLesson() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var request models.HomeWork
		err := ctx.BodyParser(&request)
		if err != nil {
			log.Println(err)
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		studentLesson, err := h.Service.GetStudentLesson(request.StudentID, request.LessonID)
		if err != nil || studentLesson.ID == 0 {
			log.Println(err)
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error":      "student has not this subject",
				"student_id": request.StudentID,
				"lesson_id":  request.LessonID,
			})
		}

		return ctx.Next()
	}
}
func (h *Handler) AdminRoleMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")
		if tokenString == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		headerParts := strings.Split(tokenString, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		claims, err := h.Service.ParseToken(headerParts[1])
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		role, err := h.Service.GetUserRole(claims.UserID)
		if err != nil {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Failed to get role"})
		}

		if role.ID != AdminRoleId {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}

		return ctx.Next()
	}
}

func (h *Handler) MarkBelongsTeacherMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		markIDParam := ctx.Params("id")
		markID, err := strconv.Atoi(markIDParam)
		if err != nil {
			log.Println(err)
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid mark ID"})
		}

		tokenString := ctx.Get("Authorization")
		if tokenString == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		headerParts := strings.Split(tokenString, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		claims, err := h.Service.ParseToken(headerParts[1])
		if err != nil {
			log.Println(err)
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		existingLesson, err := h.Service.GetLessonFromMark(markID)
		if err != nil {
			log.Println(err)
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve mark"})
		}

		if claims.UserID != existingLesson.TeacherID {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}

		return ctx.Next()
	}
}
