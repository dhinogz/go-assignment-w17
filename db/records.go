package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MongoStore) GetRecords(ctx context.Context, params RecordParams) ([]Record, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	setTotalCount := bson.D{
		{"$set", bson.D{
			{"totalCount", bson.D{
				{"$sum", "$counts"},
			}},
		}},
	}

	matchCreatedAt := bson.D{
		{"$match", bson.D{
			{"createdAt", bson.D{
				{"$gte", params.StartDate},
				{"$lte", params.EndDate},
			}},
		}},
	}

	matchTotalCount := bson.D{
		{"$match", bson.D{
			{"totalCount", bson.D{
				{"$gte", params.MinCount},
				{"$lte", params.MaxCount},
			}},
		}},
	}

	cursor, err := s.coll.Aggregate(
		ctx,
		mongo.Pipeline{setTotalCount, matchCreatedAt, matchTotalCount},
	)
	defer cursor.Close(ctx)

	var records []Record
	if err = cursor.All(ctx, &records); err != nil {
		return nil, err
	}

	return records, nil
}
