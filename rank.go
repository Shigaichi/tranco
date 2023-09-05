package tranco

import "time"

type ApiRanks struct {
	Ranks  []ApiRankEntry `json:"ranks"`
	Domain string         `json:"domain"`
}

type ApiRankEntry struct {
	Date ApiDate `json:"date"`
	Rank int     `json:"rank"`
}

type Ranks struct {
	Ranks  []RankEntry
	Domain string
}

type RankEntry struct {
	Date time.Time
	Rank int
}
