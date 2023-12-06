package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/muhammedarifp/content/commonHelp/requests"
	"github.com/muhammedarifp/content/config"
	"github.com/muhammedarifp/content/domain"
	"github.com/muhammedarifp/content/repository/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ContentUserDatabase struct {
	DB *mongo.Client
}

func NewContentUserRepo(db *mongo.Client) interfaces.ContentUserRepository {
	return &ContentUserDatabase{DB: db}
}

func (d *ContentUserDatabase) CreateNewPost(ctx context.Context, userid string, post requests.CreateNewPostRequest) (domain.Contents, error) {
	cfg := config.GetConfig()
	var emptyContentResp domain.Contents

	select {
	case <-ctx.Done():
		return emptyContentResp, errors.New("time limit reached")
	default:

	}

	_newPostId := primitive.NewObjectID()

	useridInt, _ := strconv.Atoi(userid)
	newContent := domain.Contents{
		ID:              _newPostId,
		CreateAt:        time.Now(),
		LastUpdate:      time.Now(),
		UserID:          uint(useridInt),
		ThumbnailImg:    "",
		Title:           post.Title,
		Body:            post.Body,
		Like:            0,
		IsShowReactions: post.Is_showReactions,
		IsPremium:       post.Is_premium,
		Comments:        []domain.Comment{},
		Reactions:       []domain.Reaction{},
		IsActive:        true,
	}
	_, err := d.DB.Database(cfg.DB_NAME).Collection("contents").InsertOne(context.TODO(), newContent)
	if err != nil {
		return emptyContentResp, err
	}

	// fmt.Println(res)

	return newContent, nil
}

func (d *ContentUserDatabase) CreateComment(ctx context.Context, post_id, userid string, text string) (domain.Contents, error) {
	cfg := config.GetConfig()
	var emptyContentResp domain.Contents
	select {
	case <-ctx.Done():
		return emptyContentResp, errors.New("time limit reached")
	default:

	}

	object_id, obj_err := primitive.ObjectIDFromHex(post_id)
	if obj_err != nil {
		return emptyContentResp, obj_err
	}

	filter := bson.M{"_id": object_id}
	update := bson.M{
		"$push": bson.M{"comments": domain.Comment{
			CreateAt: time.Now(),
			UserID:   userid,
			Message:  text,
		}},
	}

	_, updateErr := d.DB.Database(cfg.DB_NAME).Collection("contents").UpdateOne(ctx, filter, update)
	if updateErr != nil {
		return emptyContentResp, updateErr
	}

	return emptyContentResp, nil
}

func (d *ContentUserDatabase) LikePost(ctx context.Context, postid, userid string) (domain.Contents, error) {
	var emptyContentResp domain.Contents
	cfg := config.GetConfig()

	// Set time limit
	select {
	case <-ctx.Done():
		return emptyContentResp, errors.New("time limit reached")
	default:

	}

	// Parse postid
	objid, objidErr := primitive.ObjectIDFromHex(postid)
	if objidErr != nil {
		return emptyContentResp, errors.New("Invalid PostID")
	}

	// Update like count
	filter := bson.M{"_id": objid}
	reaction := domain.Reaction{
		CreateAt: time.Now(),
		UserID:   userid,
	}
	update := bson.M{
		"$inc":  bson.M{"like": 1},
		"$push": bson.M{"reactions": reaction},
	}
	//res, resErr := d.DB.Database(cfg.DB_NAME).Collection("contents").UpdateOne(ctx, filter, update, nil)
	opts := options.FindOneAndUpdate().SetUpsert(false).SetReturnDocument(options.After)
	res := d.DB.Database(cfg.DB_NAME).Collection("contents").FindOneAndUpdate(ctx, filter, update, opts)

	if res.Err() != nil {
		return emptyContentResp, res.Err()
	}
	var a domain.Contents
	res.Decode(&a)

	fmt.Println(a)

	return a, nil
}

func (d *ContentUserDatabase) UpdatePost(ctx context.Context, post requests.UpdatePostRequest, userID string) (domain.Contents, error) {
	cfg := config.GetConfig()

	select {
	case <-ctx.Done():
		return domain.Contents{}, errors.New("time limit reached")
	default:

	}

	// Parse ObjectID
	objID, err := primitive.ObjectIDFromHex(post.Postid)
	if err != nil {
		return domain.Contents{}, fmt.Errorf("invalid ObjectID: %w", err)
	}

	// Define filter and update
	filter := bson.M{"_id": objID, "user_id": userID}
	update := bson.M{
		"body":              post.Body,
		"is_premium":        post.Is_premium,
		"is_show_reactions": post.Is_showReactions,
		"title":             post.Title,
		"last_update":       time.Now(),
	}

	// Define options for FindOneAndUpdate
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Perform FindOneAndUpdate
	res := d.DB.Database(cfg.DB_NAME).Collection("contents").FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options)

	// Check for errors
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return domain.Contents{}, fmt.Errorf("post not found or not owned by the user")
		}
		return domain.Contents{}, fmt.Errorf("update failed: %w", res.Err())
	}

	// Decode the result into contentData
	var contentData domain.Contents
	if err := res.Decode(&contentData); err != nil {
		return domain.Contents{}, fmt.Errorf("result decoding failed: %w", err)
	}

	return contentData, nil
}

func (d *ContentUserDatabase) RemovePost(ctx context.Context, postID, userID string) (domain.Contents, error) {
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
	res := d.DB.Database(cfg.DB_NAME).Collection("contents").FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options)

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

func (d *ContentUserDatabase) GetUserPosts(ctx context.Context, userid string, page int) ([]domain.Contents, error) {
	cfg := config.GetConfig()
	collection := d.DB.Database(cfg.DB_NAME).Collection(cfg.DB_NAME)

	limit := 10
	offset := (page - 1) * limit
	options := options.Find().SetSkip(int64(offset)).SetLimit(int64(limit))
	useridInt, strconvErr := strconv.Atoi(userid)
	if strconvErr != nil {
		return []domain.Contents{}, strconvErr
	}
	filter := bson.M{
		"user_id":   useridInt,
		"is_active": true,
	}
	cursor, findErr := collection.Find(ctx, filter, options)
	if findErr != nil {
		return []domain.Contents{}, findErr
	}

	var a []domain.Contents
	if err := cursor.All(ctx, &a); err != nil {
		return []domain.Contents{}, err
	}

	return a, nil
}

func (d *ContentUserDatabase) GetallPosts(ctx context.Context, page int) ([]domain.Contents, error) {
	cfg := config.GetConfig()
	collection := d.DB.Database(cfg.DB_NAME).Collection("contents")
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
