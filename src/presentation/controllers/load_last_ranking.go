package controllers

import (
	errors "github.com/jailtonjunior94/go-clean-api/src/domain/errors"
	"github.com/jailtonjunior94/go-clean-api/src/domain/usecases"
	interfaces "github.com/jailtonjunior94/go-clean-api/src/presentation/contracts"
)

type LoadLastRankingController struct {
	LastRankingLoader usecases.LastRankingLoader
}

func NewLoadLastRankingController(u usecases.LastRankingLoader) interfaces.Controller {
	return &LoadLastRankingController{LastRankingLoader: u}
}

func (l *LoadLastRankingController) Handle() *interfaces.HttpResponse {
	rankings := l.LastRankingLoader.Load()
	if len(rankings) == 0 {
		return interfaces.ServerError(errors.RankingUnavailableError)
	}

	return interfaces.Ok(rankings)
}
