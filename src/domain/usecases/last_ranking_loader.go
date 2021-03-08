package usecases

import (
	"github.com/jailtonjunior94/go-clean-api/src/domain/entities"
)

type LastRankingLoader interface {
	Load() []entities.RankingScore
}
