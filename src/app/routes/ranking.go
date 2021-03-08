package routes

import (
	"github.com/jailtonjunior94/go-clean-api/src/app/adapters"
	"github.com/jailtonjunior94/go-clean-api/src/app/factories"

	"github.com/gofiber/fiber/v2"
)

func AddRankingRouter(router fiber.Router) {
	router.Get("/rankings/last", adapters.AdaptRoute(factories.MakeLoadLastRankingController()))
}
