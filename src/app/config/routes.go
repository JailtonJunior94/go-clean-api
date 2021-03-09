package config

import (
	"github.com/jailtonjunior94/go-clean-api/src/app/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	routes.AddRankingRouter(v1)
}
