package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grading-system-golang/internal/models"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateLesson(ctx *fiber.Ctx) error {
	var request models.Lesson
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	lessonID, err := h.Service.AddLesson(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create lesson"})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Lesson created", "lesson_id": lessonID})
}

func (h *Handler) UpdateLesson(ctx *fiber.Ctx) error {
	var request models.Lesson
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = h.Service.UpdateLesson(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update lesson"})
	}

	return ctx.JSON(fiber.Map{"message": "Lesson updated"})
}

func (h *Handler) DeleteLesson(ctx *fiber.Ctx) error {
	lessonIDParam := ctx.Params("id")
	lessonID, err := strconv.Atoi(lessonIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid lesson ID"})
	}

	err = h.Service.DeleteLesson(lessonID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete lesson"})
	}

	return ctx.JSON(fiber.Map{"message": "Lesson deleted"})
}

func (h *Handler) GetLesson(ctx *fiber.Ctx) error {
	lessonIDParam := ctx.Params("id")
	lessonID, err := strconv.Atoi(lessonIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid lesson ID"})
	}

	lesson, err := h.Service.GetLessonByID(lessonID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve lesson"})
	}

	return ctx.JSON(lesson)
}

func (h *Handler) GetLessons(ctx *fiber.Ctx) error {
	lessons, err := h.Service.GetAllLessons()
	if err != nil {
		log.Println(err) // Add this line to log the error
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve lessons"})
	}

	return ctx.JSON(lessons)
}
