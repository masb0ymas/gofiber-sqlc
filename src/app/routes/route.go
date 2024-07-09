package routes

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/masb0ymas/go-utils/pkg"
)

func InitializeRoutes(app *fiber.App) {
	// base route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"code":      http.StatusOK,
			"message":   "gofiber-sqlc",
			"maintaner": "masb0ymas, <n.fajri@outlook.com>",
			"source":    "https://github.com/masb0ymas/gofiber-sqlc",
		},
		)
	})

	// health route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"code":    http.StatusOK,
			"cpu":     runtime.NumCPU(),
			"date":    pkg.TimeIn("ID").Format(time.RFC850),
			"golang":  runtime.Version(),
			"gofiber": fiber.Version,
			"status":  "Ok",
		},
		)
	})

	// forbidden route version
	app.Get("/v1", func(c *fiber.Ctx) error {
		return c.Status(http.StatusForbidden).JSON(fiber.NewError(http.StatusForbidden))
	})

	// initialize route v1
	RouteV1(app)

	// not found route
	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(fiber.NewError(http.StatusNotFound, "Sorry, HTTP resource you are looking for was not found."))
	})
}
