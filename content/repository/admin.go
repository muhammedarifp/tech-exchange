package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/muhammedarifp/content/config"
	"github.com/muhammedarifp/content/domain"
	"github.com/muhammedarifp/content/repository/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AdminContentRepository struct {
	db *mongo.Client
}

func NewAdminContentRepository(db *mongo.Client) interfaces.AdminContentRepo {
	return &AdminContentRepository{db: db}
}

func (d *AdminContentRepository) GetallPosts(ctx context.Context, page int) ([]domain.Contents, error) {
	cfg := config.GetConfig()
	collection := d.db.Database(cfg.DB_NAME).Collection("contents")
	limit := 10
	offset := (page - 1) * limit
	options := options.Find().SetSkip(int64(offset)).SetLimit(10)

	//filter
	filter := bson.M{"is_active": true}
	cursor, findErr := collection.Find(ctx, filter, options)
	if findErr != nil {
		return []domain.Contents{}, fmt.Errorf("find error %w", findErr)
	}

	var s []domain.Contents
	if err := cursor.All(ctx, &s); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(s)

	return s, nil
}
