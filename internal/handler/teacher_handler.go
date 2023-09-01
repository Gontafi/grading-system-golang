package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grading-system-golang/internal/models"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateTeacher(ctx *fiber.Ctx) error {
	var request models.Teacher
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	teacherID, err := h.Service.AddTeacher(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create teacher"})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Teacher created", "teacher_id": teacherID})
}

func (h *Handler) UpdateTeacher(ctx *fiber.Ctx) error {
	var request models.Teacher
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = h.Service.UpdateTeacher(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update teacher"})
	}

	return ctx.JSON(fiber.Map{"message": "Teacher updated"})
}

func (h *Handler) DeleteTeacher(ctx *fiber.Ctx) error {
	teacherIDParam := ctx.Params("id")
	teacherID, err := strconv.Atoi(teacherIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid teacher ID"})
	}

	err = h.Service.DeleteTeacher(teacherID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete teacher"})
	}

	return ctx.JSON(fiber.Map{"message": "Teacher deleted"})
}

func (h *Handler) GetTeacher(ctx *fiber.Ctx) error {
	teacherIDParam := ctx.Params("id")
	teacherID, err := strconv.Atoi(teacherIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid teacher ID"})
	}

	teacher, err := h.Service.GetTeacherByID(teacherID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve teacher"})
	}

	return ctx.JSON(teacher)
}

func (h *Handler) GetTeachers(ctx *fiber.Ctx) error {
	teachers, err := h.Service.GetAllTeachers()
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve teachers"})
	}

	return ctx.JSON(teachers)
}
