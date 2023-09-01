package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) GetTopRatingFromCache(ctx *fiber.Ctx) error {
	periodParam := ctx.Query("period")
	limitParam := ctx.Query("limit")

	period, err := time.ParseDuration(periodParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid period"})
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit"})
	}

	ratings, err := h.Service.GetTopRatingFromCache(period, limit)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve top ratings", "err": err})
	}

	return ctx.JSON(ratings)
}

func (h *Handler) GetTopRatingByLessonFromCache(ctx *fiber.Ctx) error {
	lessonIDParam := ctx.Params("lessonID")
	periodParam := ctx.Query("period")
	limitParam := ctx.Query("limit")

	lessonID, err := strconv.Atoi(lessonIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid lesson ID"})
	}

	period, err := time.ParseDuration(periodParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid period"})
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit"})
	}

	ratings, err := h.Service.GetTopRatingByLessonFromCache(lessonID, period, limit)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve top ratings by lesson", "err": err})
	}

	return ctx.JSON(ratings)
}
