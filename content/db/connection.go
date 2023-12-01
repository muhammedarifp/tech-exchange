package db

import (
	"context"

	"github.com/muhammedarifp/content/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDbConnection() (*mongo.Client, error) {
	cfg := config.GetConfig()
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MONGO_PORT))
	if err != nil {
		return nil, err
	}

	return client, nil
}
