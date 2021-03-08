package factories

import (
	"github.com/jailtonjunior94/go-clean-api/src/application/services"
	"github.com/jailtonjunior94/go-clean-api/src/infrastructure/repositories"
	"github.com/jailtonjunior94/go-clean-api/src/presentation/contracts"
	"github.com/jailtonjunior94/go-clean-api/src/presentation/controllers"
)

func MakeLoadLastRankingController() contracts.Controller {
	repository := repositories.NewFakeRankingRepository()
	loader := services.NewLastRankingLoaderService(repository)
	return controllers.NewLoadLastRankingController(loader)
}
