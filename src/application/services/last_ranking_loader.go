package services

import (
	"fmt"
	"time"

	"github.com/jailtonjunior94/go-clean-api/src/application/contracts"
	"github.com/jailtonjunior94/go-clean-api/src/domain/entities"
	errors "github.com/jailtonjunior94/go-clean-api/src/domain/errors"
	"github.com/jailtonjunior94/go-clean-api/src/domain/usecases"
)

type LastRankingLoaderService struct {
	repository contracts.LoadLastRankingRepository
}

func NewLastRankingLoaderService(r contracts.LoadLastRankingRepository) usecases.LastRankingLoader {
	return &LastRankingLoaderService{repository: r}
}

func (l *LastRankingLoaderService) Load() []entities.RankingScore {
	if time.Now().Hour() > 22 {
		fmt.Println(errors.RankingUnavailableError)
	}

	rankingScore := make([]entities.RankingScore, 0)
	lastRanking := l.repository.LoadLastRanking()
	for _, l := range lastRanking {
		ranking := entities.RankingScore{
			Player: entities.Player{
				Name:    l.Player.Name,
				Country: l.Player.Country,
			},
			Score:     l.Score,
			MatchDate: l.MatchDate,
		}

		rankingScore = append(rankingScore, ranking)
	}

	return rankingScore
}
