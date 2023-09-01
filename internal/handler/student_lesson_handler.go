package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) AddStudentToLesson(ctx *fiber.Ctx) error {
	studentIDParam := ctx.Params("student_id")
	studentID, err := strconv.Atoi(studentIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid student ID"})
	}

	lessonIDParam := ctx.Params("lesson_id")
	lessonID, err := strconv.Atoi(lessonIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid lesson ID"})
	}

	id, err := h.Service.AddStudentToLesson(studentID, lessonID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add student to lesson"})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Student added to lesson", "id": id})
}

func (h *Handler) RemoveStudentFromLesson(ctx *fiber.Ctx) error {
	studentIDParam := ctx.Params("student_id")
	studentID, err := strconv.Atoi(studentIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid student ID"})
	}

	lessonIDParam := ctx.Params("lesson_id")
	lessonID, err := strconv.Atoi(lessonIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid lesson ID"})
	}

	err = h.Service.RemoveStudentFromLesson(studentID, lessonID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to remove student from lesson"})
	}

	return ctx.JSON(fiber.Map{"message": "Student removed from lesson"})
}

func (h *Handler) GetStudentsForLesson(ctx *fiber.Ctx) error {
	lessonIDParam := ctx.Params("lesson_id")
	lessonID, err := strconv.Atoi(lessonIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid lesson ID"})
	}

	students, err := h.Service.GetStudentsForLesson(lessonID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve students for lesson"})
	}

	return ctx.JSON(students)
}

func (h *Handler) GetLessonsForStudent(ctx *fiber.Ctx) error {
	studentIDParam := ctx.Params("student_id")
	studentID, err := strconv.Atoi(studentIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid student ID"})
	}

	lessons, err := h.Service.GetLessonsForStudent(studentID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve lessons for student"})
	}

	return ctx.JSON(lessons)
}

func (h *Handler) GetStudentLesson(ctx *fiber.Ctx) error {
	studentIDParam := ctx.Params("student_id")
	lessonIDParam := ctx.Params("lesson_id")
	studentID, err := strconv.Atoi(studentIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid student ID"})
	}
	lessonID, err := strconv.Atoi(lessonIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid lesson ID"})
	}

	lesson, err := h.Service.GetStudentLesson(studentID, lessonID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve lessons"})
	}

	return ctx.JSON(lesson)
}
