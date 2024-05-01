package db

import (
	"time"
)

type RecordParams struct {
	StartDate time.Time
	EndDate   time.Time
	MinCount  int
	MaxCount  int
}

type Record struct {
	Key        string    `bson:"_id,omitempty"       json:"key"`
	CreatedAt  time.Time `bson:"createdAt,omitempty" json:"createdAt"`
	Counts     []int     `bson:"counts"              json:"-"`
	TotalCount int       `bson:"totalCount"          json:"totalCount"`
}
