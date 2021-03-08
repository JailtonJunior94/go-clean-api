package adapters

import (
	"github.com/jailtonjunior94/go-clean-api/src/presentation/contracts"

	"github.com/gofiber/fiber/v2"
)

func AdaptRoute(controller contracts.Controller) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		httpResponse := controller.Handle()
		return c.Status(httpResponse.StatusCode).JSON(httpResponse.Data)
	}
}
