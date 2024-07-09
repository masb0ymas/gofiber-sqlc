package controller

import (
	"context"
	"gofiber-sqlc/src/app/schema"
	"gofiber-sqlc/src/app/service"
	"gofiber-sqlc/src/pkg/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type NewRoleParams struct {
	Name string
}

func GetRoles(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	roleService := service.NewRoleService()
	roles, total, err := roleService.FindAll(ctx, c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been received",
		"data":    roles,
		"total":   total,
	})
}

func GetRole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	roleService := service.NewRoleService()
	role, err := roleService.FindOne(ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been received",
		"data":    role,
	})
}

func NewRoles(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	formData := new(schema.RoleSchema)
	if code, message, errors := utils.ParseFormDataAndValidate(c, formData); errors != nil {
		return c.Status(int(code)).JSON(fiber.Map{
			"code":    int(code),
			"message": message,
			"errors":  errors,
		})
	}

	roleService := service.NewRoleService()
	data, err := roleService.Create(ctx, formData.Name)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    http.StatusCreated,
		"message": "data has been created",
		"data":    data,
	})
}
