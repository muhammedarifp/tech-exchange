package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/muhammedarifp/content/commonHelp/requests"
	"github.com/muhammedarifp/content/config"
	"github.com/muhammedarifp/content/domain"
	"github.com/muhammedarifp/content/repository/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	defer d.DB.Disconnect(ctx)

	select {
	case <-ctx.Done():
		return emptyContentResp, errors.New("time limit reached")
	default:
		time.Sleep(time.Second)
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
	}
	res, err := d.DB.Database(cfg.DB_NAME).Collection("contents").InsertOne(context.TODO(), newContent)
	if err != nil {
		return emptyContentResp, err
	}

	fmt.Println(res)

	return newContent, nil
}

func (d *ContentUserDatabase) CreateComment(ctx context.Context, post_id, userid string, text string) (domain.Contents, error) {
	cfg := config.GetConfig()
	var emptyContentResp domain.Contents
	select {
	case <-ctx.Done():
		return emptyContentResp, errors.New("time limit reached")
	default:
		time.Sleep(time.Second)
	}

	object_id, obj_err := primitive.ObjectIDFromHex(post_id)
	if obj_err != nil {
		return emptyContentResp, obj_err
	}

	filter := bson.M{"_id": object_id}
	fmt.Println(filter)
	update := bson.M{
		"$push": bson.M{"comments": domain.Comment{
			CreateAt: time.Now(),
			UserID:   userid,
			Message:  text,
		}},
	}

	updateRes, updateErr := d.DB.Database(cfg.DB_NAME).Collection("contents").UpdateOne(ctx, filter, update)
	if updateErr != nil {
		fmt.Println("-->", updateErr.Error())
		return emptyContentResp, updateErr
	}
	fmt.Println(updateRes.MatchedCount)

	return emptyContentResp, nil
}

func (d *ContentUserDatabase) LikePost(ctx context.Context, postid string) (domain.Contents, error) {
	var emptyContentResp domain.Contents
	cfg := config.GetConfig()
	select {
	case <-ctx.Done():
		return emptyContentResp, errors.New("time limit reached")
	default:
		time.Sleep(time.Second)
	}

	objid, objidErr := primitive.ObjectIDFromHex(postid)
	if objidErr != nil {
		return emptyContentResp, errors.New("Postid is invalid")
	}
	filter := bson.M{"_id": objid}
	update := bson.M{"$inc": bson.M{"like": 1}}
	res, resErr := d.DB.Database(cfg.DB_NAME).Collection("contents").UpdateOne(ctx, filter, update, nil)
	if resErr != nil || res.MatchedCount <= 0 {
		return emptyContentResp, resErr
	}

	fmt.Println(res.MatchedCount)

	return emptyContentResp, nil
}
