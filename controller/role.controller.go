package controller

import (
	"context"
	"fmt"
	"gofiber-sqlc/database/schema"
	"gofiber-sqlc/pkg/utils"
	"gofiber-sqlc/service"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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
		"message": fmt.Sprintf("%d data has been received", total),
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

func NewRole(c *fiber.Ctx) error {
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

func UpdateRole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	formData := new(schema.RoleSchema)
	if code, message, errors := utils.ParseFormDataAndValidate(c, formData); errors != nil {
		return c.Status(int(code)).JSON(fiber.Map{
			"code":    int(code),
			"message": message,
			"errors":  errors,
		})
	}

	roleService := service.NewRoleService()
	data, err := roleService.Update(ctx, id, formData.Name)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been updated",
		"data":    data,
	})
}

func RestoreRole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	roleService := service.NewRoleService()
	err = roleService.Restore(ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been restored",
	})
}

func SoftDeleteRole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	roleService := service.NewRoleService()
	err = roleService.SoftDelete(ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been deleted",
	})
}

func ForceDeleteRole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	roleService := service.NewRoleService()
	err = roleService.ForceDelete(ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been deleted",
	})
}
