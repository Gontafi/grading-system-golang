package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/grading-system-golang/internal/models"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateMark(ctx *fiber.Ctx) error {
	var request models.Mark
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	tokenString := ctx.Get("Authorization")
	claims, err := ParseToken(tokenString)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	if claims.Role != "teacher" || claims.UserID != request.TeacherID {
		log.Println(errors.New("Access denied")) // Log a custom error message
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}

	_, err = h.Service.GetStudentLesson(request.StudentID, request.LessonID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":      "student has not this subject",
			"student_id": request.StudentID,
			"lesson_id":  request.LessonID,
		})
	}

	markID, err := h.Service.CreateMark(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create mark", "err": err})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Mark created", "mark_id": markID})
}

func (h *Handler) UpdateMark(ctx *fiber.Ctx) error {

	markIDParam := ctx.Params("id")
	markID, err := strconv.Atoi(markIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid mark ID"})
	}

	var request models.Mark
	err = ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	tokenString := ctx.Get("Authorization")
	claims, err := ParseToken(tokenString)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	existingMark, err := h.Service.GetMarkByID(markID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve mark"})
	}

	if claims.Role != "teacher" || claims.UserID != existingMark.TeacherID {
		log.Println(errors.New("Access denied"))
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}

	err = h.Service.UpdateMark(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update mark"})
	}

	return ctx.JSON(fiber.Map{"message": "Mark updated"})
}

func (h *Handler) DeleteMark(ctx *fiber.Ctx) error {

	markIDParam := ctx.Params("id")
	markID, err := strconv.Atoi(markIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid mark ID"})
	}

	tokenString := ctx.Get("Authorization")
	claims, err := ParseToken(tokenString)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	existingMark, err := h.Service.GetMarkByID(markID)
	if err != nil {
		log.Println(err) // Add this line to log the error
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve mark"})
	}

	if claims.Role != "teacher" || claims.UserID != existingMark.TeacherID {
		log.Println(errors.New("Access denied"))
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}

	err = h.Service.DeleteMark(markID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete mark"})
	}

	return ctx.JSON(fiber.Map{"message": "Mark deleted"})
}

func (h *Handler) GetMark(ctx *fiber.Ctx) error {
	markIDParam := ctx.Params("id")
	markID, err := strconv.Atoi(markIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid mark ID"})
	}

	mark, err := h.Service.GetMarkByID(markID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve mark"})
	}

	return ctx.JSON(mark)
}

func (h *Handler) GetMarks(ctx *fiber.Ctx) error {
	marks, err := h.Service.GetAllMarks()
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve marks"})
	}

	return ctx.JSON(marks)
}
