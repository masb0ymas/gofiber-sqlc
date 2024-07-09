package main

import (
	"gofiber-sqlc/src/app/routes"
	"gofiber-sqlc/src/database"
	"gofiber-sqlc/src/pkg/config"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func init() {
	database.ConnectDB()
}

func main() {
	// app config
	app := fiber.New()
	port := config.Env("APP_PORT", "8000")
	envRateLimit := config.Env("APP_RATE_LIMIT", "100")
	rateLimit, _ := strconv.Atoi(envRateLimit)

	// use middleware
	app.Use(cors.New(config.Cors()))
	app.Use(compress.New())
	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{Max: rateLimit}))
	app.Use(requestid.New())
	app.Use(recover.New())

	// static file
	app.Static("/", "./public")

	// initialize routes
	routes.InitializeRoutes(app)

	// listen app
	log.Fatal(app.Listen(":" + port))
}
