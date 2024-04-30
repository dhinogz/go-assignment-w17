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

type Record struct {
	Key        string    `bson:"_id,omitempty"       json:"key"`
	CreatedAt  time.Time `bson:"createdAt,omitempty" json:"createdAt"`
	Counts     []int     `bson:"counts"              json:"-"`
	TotalCount int       `bson:"totalCount"          json:"totalCount"`
}
