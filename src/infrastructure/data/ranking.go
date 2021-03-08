package data

import (
	"time"

	"github.com/jailtonjunior94/go-clean-api/src/application/dtos"
)

func NewRankingData() []dtos.RankingScore {
	heroes := make([]dtos.Hero, 0)
	rankings := make([]dtos.RankingScore, 0)

	hero := dtos.Hero{
		Name:  "Fee",
		Level: 30,
	}

	heroes = append(heroes, hero)

	rankingOne := dtos.RankingScore{
		Player: dtos.Player{
			Name:    "Rodrigo Manguinho",
			Country: "Brazil",
		},
		MatchDate: time.Now(),
		Score:     1578625200000,
		Heroes:    heroes,
	}

	rankingTwo := dtos.RankingScore{
		Player: dtos.Player{
			Name:    "Rodrigo Branas",
			Country: "Brazil",
		},
		MatchDate: time.Now(),
		Score:     1578625200000,
		Heroes:    heroes,
	}

	rankings = append(rankings, rankingOne)
	rankings = append(rankings, rankingTwo)

	return rankings
}
