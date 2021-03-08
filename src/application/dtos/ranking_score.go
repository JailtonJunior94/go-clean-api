package dtos

import "time"

type RankingScore struct {
	Player    Player
	Score     int64
	MatchDate time.Time
	Heroes    []Hero
}

type Player struct {
	Name    string
	Country string
}

type Hero struct {
	Name  string
	Level int32
}
