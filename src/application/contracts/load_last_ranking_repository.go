package contracts

import "github.com/jailtonjunior94/go-clean-api/src/application/dtos"

type LoadLastRankingRepository interface {
	LoadLastRanking() []dtos.RankingScore
}
