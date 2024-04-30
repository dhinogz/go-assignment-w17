package db

import (
	"context"
	"time"
)

// TODO: actually call mongo database to fetch record
func (s *MongoStore) GetRecords(ctx context.Context, params RecordParams) ([]Record, error) {
	return []Record{
		{
			Key:        "1",
			CreatedAt:  time.Now(),
			TotalCount: 1,
		},
		{
			Key:        "2",
			CreatedAt:  time.Now(),
			TotalCount: 2,
		},
	}, nil
}
