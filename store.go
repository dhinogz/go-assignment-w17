package main

import (
	"context"
	"errors"
	"time"
	//
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: Fill in this data with env variables
const DBNAME = "records"
const userColl = "records"

type Store interface {
	GetRecords(context.Context, RecordParams) Response
}

type MongoStore struct {
	coll   *mongo.Collection
	client *mongo.Client
}

func NewMongoStore(client *mongo.Client) (*MongoStore, error) {
	if client == nil {
		return nil, errors.New("new mongo store: client provided is nil")
	}
	return &MongoStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(userColl),
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
