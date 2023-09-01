package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grading-system-golang/internal/models"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateStudent(ctx *fiber.Ctx) error {
	var request models.Student
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	studentID, err := h.Service.AddStudent(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create student"})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Student created", "student_id": studentID})
}

func (h *Handler) UpdateStudent(ctx *fiber.Ctx) error {
	var request models.Student
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = h.Service.UpdateStudent(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update student"})
	}

	return ctx.JSON(fiber.Map{"message": "Student updated"})
}

func (h *Handler) DeleteStudent(ctx *fiber.Ctx) error {
	studentIDParam := ctx.Params("id")
	studentID, err := strconv.Atoi(studentIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid student ID"})
	}

	err = h.Service.DeleteStudent(studentID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete student"})
	}

	return ctx.JSON(fiber.Map{"message": "Student deleted"})
}

func (h *Handler) GetStudent(ctx *fiber.Ctx) error {
	studentIDParam := ctx.Params("id")
	studentID, err := strconv.Atoi(studentIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid student ID"})
	}

	student, err := h.Service.GetStudentByID(studentID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve student"})
	}

	return ctx.JSON(student)
}

func (h *Handler) GetStudents(ctx *fiber.Ctx) error {
	students, err := h.Service.GetAllStudents()
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve students"})
	}

	return ctx.JSON(students)
}
