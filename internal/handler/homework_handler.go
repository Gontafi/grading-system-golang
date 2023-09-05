package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grading-system-golang/internal/models"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateHomeWork(ctx *fiber.Ctx) error {
	var request models.HomeWork
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	homeWorkID, err := h.Service.AddHomeWork(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create homework"})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Homework created", "homework_id": homeWorkID})
}

func (h *Handler) UpdateHomeWork(ctx *fiber.Ctx) error {
	homeWorkID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid homework ID"})
	}

	var request models.HomeWork
	err = ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	request.ID = homeWorkID

	err = h.Service.UpdateHomeWork(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update homework"})
	}

	return ctx.JSON(fiber.Map{"message": "Homework updated"})
}

func (h *Handler) DeleteHomeWork(ctx *fiber.Ctx) error {
	homeWorkIDParam := ctx.Params("id")
	homeWorkID, err := strconv.Atoi(homeWorkIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid homework ID"})
	}

	err = h.Service.DeleteHomeWork(homeWorkID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete homework"})
	}

	return ctx.JSON(fiber.Map{"message": "Homework deleted"})
}

func (h *Handler) GetHomeWork(ctx *fiber.Ctx) error {
	homeWorkIDParam := ctx.Params("id")
	homeWorkID, err := strconv.Atoi(homeWorkIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid homework ID"})
	}

	homeWork, err := h.Service.GetHomeWorkById(homeWorkID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve homework"})
	}

	return ctx.JSON(homeWork)
}

func (h *Handler) GetHomeWorks(ctx *fiber.Ctx) error {
	homeWorks, err := h.Service.GetAllHomeWorks()
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve homeworks"})
	}

	return ctx.JSON(homeWorks)
}
