package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grading-system-golang/internal/models"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateUser(ctx *fiber.Ctx) error {
	var request models.User
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	userID, err := h.Service.AddUser(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "User created", "user_id": userID})
}

func (h *Handler) UpdateUser(ctx *fiber.Ctx) error {
	var request models.User
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = h.Service.UpdateUser(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return ctx.JSON(fiber.Map{"message": "User updated"})
}

func (h *Handler) DeleteUser(ctx *fiber.Ctx) error {
	userIDParam := ctx.Params("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	err = h.Service.DeleteUser(userID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return ctx.JSON(fiber.Map{"message": "User deleted"})
}

func (h *Handler) GetUser(ctx *fiber.Ctx) error {
	userIDParam := ctx.Params("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := h.Service.GetUserByID(userID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user"})
	}

	return ctx.JSON(user)
}

func (h *Handler) GetUsers(ctx *fiber.Ctx) error {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve users"})
	}

	return ctx.JSON(users)
}
