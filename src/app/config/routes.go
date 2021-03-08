package config

import (
	"github.com/jailtonjunior94/go-clean-api/src/app/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Get("/me", me)

	routes.AddRankingRouter(v1)
}

func me(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Me{Message: "I'm a project, Go Challenge Appointments"})

}

type Me struct {
	Message string `json:"message"`
}
