package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// TODO: read https://www.mongodb.com/docs/drivers/go/master/fundamentals/aggregation/
func (s *MongoStore) GetRecords(ctx context.Context, params RecordParams) ([]Record, error) {
	// TODO: Add filter here
	cursor, err := s.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var records []Record
	if err = cursor.All(ctx, &records); err != nil {
		return nil, err
	}

	// TODO: Use Mongo aggregation to get total count of records
	for i := 0; i < len(records); i++ {
		for _, c := range records[i].Counts {
			records[i].TotalCount += c
		}
	}

	return records, nil
}
