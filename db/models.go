package db

import (
	"time"
)

type RecordParams struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

// TODO: add bson fields
type Record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}
