package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) AddRoleUser(ctx *fiber.Ctx) error {
	roleIDParam := ctx.Params("role_id")
	roleID, err := strconv.Atoi(roleIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role ID"})
	}

	userIDParam := ctx.Params("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	id, err := h.Service.AddRoleUser(roleID, userID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add role user"})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Role user added", "id": id})
}

func (h *Handler) RemoveRoleUser(ctx *fiber.Ctx) error {
	roleIDParam := ctx.Params("role_id")
	roleID, err := strconv.Atoi(roleIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role ID"})
	}

	userIDParam := ctx.Params("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	err = h.Service.RemoveRoleUser(roleID, userID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to remove role user"})
	}

	return ctx.JSON(fiber.Map{"message": "Role user removed"})
}

func (h *Handler) GetUsersForRole(ctx *fiber.Ctx) error {
	roleIDParam := ctx.Params("role_id")
	roleID, err := strconv.Atoi(roleIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role ID"})
	}

	users, err := h.Service.GetUsersForRole(roleID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve users for role"})
	}

	return ctx.JSON(users)
}

func (h *Handler) GetRolesForUser(ctx *fiber.Ctx) error {
	userIDParam := ctx.Params("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	roles, err := h.Service.GetRolesForUser(userID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve roles for user"})
	}

	return ctx.JSON(roles)
}

func (h *Handler) GetUserRole(ctx *fiber.Ctx) error {
	userIDParam := ctx.Params("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	role, err := h.Service.GetUserRole(userID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve role user"})
	}

	return ctx.JSON(role)
}
