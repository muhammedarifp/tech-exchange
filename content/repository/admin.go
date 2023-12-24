package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/muhammedarifp/content/config"
	"github.com/muhammedarifp/content/domain"
	"github.com/muhammedarifp/content/repository/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (d *AdminContentRepository) RemovePost(ctx context.Context, postID, userID string) (domain.Contents, error) {
	cfg := config.GetConfig()

	// Parse ObjectID
	objID, objIDErr := primitive.ObjectIDFromHex(postID)
	if objIDErr != nil {
		return domain.Contents{}, fmt.Errorf("invalid post ID: %w", objIDErr)
	}

	// Define filter and update
	useridInt, strconvErr := strconv.Atoi(userID)
	if strconvErr != nil {
		return domain.Contents{}, fmt.Errorf("string convert error : %w", strconvErr)
	}
	filter := bson.M{"_id": objID, "user_id": useridInt}
	update := bson.M{"is_active": false}

	// Define options for FindOneAndUpdate
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Perform FindOneAndUpdate
	res := d.db.Database(cfg.DB_NAME).Collection("contents").FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options)

	// Check for errors
	if res.Err() != nil {
		return domain.Contents{}, fmt.Errorf("update failed: %w", res.Err())
	}

	// Decode the result into content
	var content domain.Contents
	if err := res.Decode(&content); err != nil {
		return domain.Contents{}, fmt.Errorf("result decoding failed: %w", err)
	}

	return content, nil
}

func (d *AdminContentRepository) AddNewTag(ctx context.Context, tag string) (domain.Tags, error) {
	cfg := config.GetConfig()

	newTag := domain.Tags{
		ID:       primitive.NewObjectID(),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		Tag:      tag,
		Posts:    []primitive.ObjectID{},
		IsActive: true,
	}

	_, err := d.db.Database(cfg.DB_NAME).Collection("tags").InsertOne(ctx, newTag)
	if err != nil {
		log.Fatal(err.Error())
		return domain.Tags{}, err
	}

	return newTag, nil
}
