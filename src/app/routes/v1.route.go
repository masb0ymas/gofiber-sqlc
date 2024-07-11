package routes

import (
	"gofiber-sqlc/src/app/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteV1(app *fiber.App) {
	// group v1
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")

		return c.Next()
	})

	roleHandler := v1.Group("/role")
	roleHandler.Get("/", controller.GetRoles)
	roleHandler.Get("/:id", controller.GetRole)
	roleHandler.Post("/", controller.NewRole)
	roleHandler.Put("/:id", controller.UpdateRole)
	roleHandler.Put("/restore/:id", controller.RestoreRole)
	roleHandler.Put("/soft-delete/:id", controller.SoftDeleteRole)
	roleHandler.Delete("/force-delete/:id", controller.ForceDeleteRole)

	userHandler := v1.Group("/user")
	userHandler.Get("/", controller.GetUsers)
	userHandler.Get("/:id", controller.GetUser)
	userHandler.Post("/", controller.NewUser)
}
