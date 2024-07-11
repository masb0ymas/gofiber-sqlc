package controller

import (
	"context"
	"fmt"
	"gofiber-sqlc/src/app/service"
	"gofiber-sqlc/src/database/sqlc"
	"gofiber-sqlc/src/pkg/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/masb0ymas/go-utils/argon2"
)

func GetUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userService := service.NewUserService()
	users, total, err := userService.FindAll(ctx, c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": fmt.Sprintf("%d data has been received", total),
		"data":    users,
		"total":   total,
	})
}

func GetUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	userService := service.NewUserService()
	user, err := userService.FindOne(ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been received",
		"data":    user,
	})
}

func NewUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	formData := new(sqlc.NewUserParams)
	if code, message, errors := utils.ParseFormDataAndValidate(c, formData); errors != nil {
		return c.Status(int(code)).JSON(fiber.Map{
			"code":    int(code),
			"message": message,
			"errors":  errors,
		})
	}

	userService := service.NewUserService()
	data, err := userService.Create(ctx, sqlc.NewUserParams{
		Fullname:    formData.Fullname,
		Email:       formData.Email,
		Password:    argon2.Generate(formData.Password),
		Phone:       formData.Phone,
		TokenVerify: formData.TokenVerify,
		Address:     formData.Address,
		IsActive:    formData.IsActive,
		IsBlocked:   formData.IsBlocked,
		RoleID:      formData.RoleID,
	})

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
