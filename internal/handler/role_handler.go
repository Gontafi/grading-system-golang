package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grading-system-golang/internal/models"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateRole(ctx *fiber.Ctx) error {
	var request models.Role
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	roleID, err := h.Service.AddRole(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create role"})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Role created", "role_id": roleID})
}

func (h *Handler) UpdateRole(ctx *fiber.Ctx) error {
	var request models.Role
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = h.Service.UpdateRole(request)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update role"})
	}

	return ctx.JSON(fiber.Map{"message": "Role updated"})
}

func (h *Handler) DeleteRole(ctx *fiber.Ctx) error {
	roleIDParam := ctx.Params("id")
	roleID, err := strconv.Atoi(roleIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role ID"})
	}

	err = h.Service.DeleteRole(roleID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete role"})
	}

	return ctx.JSON(fiber.Map{"message": "Role deleted"})
}

func (h *Handler) GetRole(ctx *fiber.Ctx) error {
	roleIDParam := ctx.Params("id")
	roleID, err := strconv.Atoi(roleIDParam)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role ID"})
	}

	role, err := h.Service.GetRoleByID(roleID)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve role"})
	}

	return ctx.JSON(role)
}

func (h *Handler) GetRoles(ctx *fiber.Ctx) error {
	roles, err := h.Service.GetAllRoles()
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve roles"})
	}

	return ctx.JSON(roles)
}
