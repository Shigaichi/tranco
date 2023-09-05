package tranco

import "time"

type ApiListMetadata struct {
	ListId        string           `json:"list_id"`
	Available     bool             `json:"available"`
	Failed        bool             `json:"failed"`
	Download      string           `json:"download"`
	CreatedOn     ApiTimestamp     `json:"created_on"`
	Configuration ApiConfiguration `json:"configuration"`
}

type ApiConfiguration struct {
	Providers         []string `json:"providers"`
	ListPrefix        string   `json:"listPrefix"`
	EndDate           ApiDate  `json:"endDate"`
	FilterTLD         string   `json:"filterTLD"`
	FilterPLD         string   `json:"filterPLD"`
	CombinationMethod string   `json:"combinationMethod"`
	StartDate         ApiDate  `json:"startDate"`
	IsDailyList       bool     `json:"isDailyList"`
}

type ListMetadata struct {
	ListId        string        `json:"list_id"`
	Available     bool          `json:"available"`
	Failed        bool          `json:"failed"`
	Download      string        `json:"download"`
	CreatedOn     time.Time     `json:"created_on"`
	Configuration Configuration `json:"configuration"`
}

type Configuration struct {
	Providers         []string  `json:"providers"`
	ListPrefix        string    `json:"listPrefix"`
	EndDate           time.Time `json:"endDate"`
	FilterTLD         string    `json:"filterTLD"`
	FilterPLD         string    `json:"filterPLD"`
	CombinationMethod string    `json:"combinationMethod"`
	StartDate         time.Time `json:"startDate"`
	IsDailyList       bool      `json:"isDailyList"`
}
