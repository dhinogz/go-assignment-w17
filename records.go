package main

import (
	"time"
)

type RecordParams struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	MinCount  int       `json:"minCount"`
	MaxCount  int       `json:"maxCount"`
}

type Response struct {
	Code    int              `json:"code"`
	Msg     string           `json:"msg"`
	Records []RecordResponse `json:"records"`
}

type RecordResponse struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}
