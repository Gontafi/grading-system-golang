package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grading-system-golang/internal/models"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CheckHomeWorkAndPutGrades(ctx *fiber.Ctx) error {
	var request models.Mark
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format"})
	}

	markID, err := h.Service.CreateMark(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create mark", "err": err})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Mark created", "mark_id": markID})
}

func (h *Handler) UpdateMark(ctx *fiber.Ctx) error {

	var request models.Mark

	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
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
