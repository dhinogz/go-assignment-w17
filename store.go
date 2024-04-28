package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName     = "getircase-study"
	recordColl = "records"
)

type Store interface {
	GetRecords(context.Context, RecordParams) Response
}

type MongoStore struct {
	coll   *mongo.Collection
	client *mongo.Client
}

func NewMongoStore(mongoURI string) (*MongoStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("could not initialize new mongo client: %+v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("could not ping mongo database: %+v", err)
	}
	slog.Info("Connected to database!")
	return &MongoStore{
		client: client,
		coll:   client.Database(dbName).Collection(recordColl),
	}, nil
}

func (s *MongoStore) GetRecords(ctx context.Context, params RecordParams) Response {
	return Response{
		Code: 0,
		Msg:  "",
		Records: []RecordResponse{
			{
				Key:        "",
				CreatedAt:  time.Now(),
				TotalCount: 0,
			},
		},
	}
}
