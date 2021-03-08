package repositories

import (
	"github.com/jailtonjunior94/go-clean-api/src/application/contracts"
	"github.com/jailtonjunior94/go-clean-api/src/application/dtos"
	"github.com/jailtonjunior94/go-clean-api/src/infrastructure/data"
)

type FakeRankingRepository struct {
}

func NewFakeRankingRepository() contracts.LoadLastRankingRepository {
	return &FakeRankingRepository{}
}

func (f *FakeRankingRepository) LoadLastRanking() []dtos.RankingScore {
	return data.NewRankingData()
}
