package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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

	newObjArr := make([]primitive.ObjectID, len(post.Labels))
	for i := 0; i < len(post.Labels); i++ {
		objid, objConvErr := primitive.ObjectIDFromHex(post.Labels[i])
		if objConvErr != nil {
			return domain.Contents{}, fmt.Errorf("invalid format for label '%s': %w", post.Labels[i], objConvErr)
		}
		newObjArr = append(newObjArr, objid)
	}

	_newPostId := primitive.NewObjectID()

	filter := bson.M{"_id": bson.M{"$in": newObjArr}}
	update := bson.M{"$push": bson.M{"posts": _newPostId}}

	_, updateErr := d.DB.Database(cfg.DB_NAME).Collection("tags").UpdateMany(ctx, filter, update, nil)
	if updateErr != nil {
		return domain.Contents{}, updateErr
	}

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
		Labels:          post.Labels,
		IsActive:        true,
	}
	_, err := d.DB.Database(cfg.DB_NAME).Collection("contents").InsertOne(context.TODO(), newContent)
	if err != nil {
		return emptyContentResp, err
	}

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

func (d *ContentUserDatabase) FollowTag(ctx context.Context, userid string, req requests.FollowTagReq) (domain.Interests, error) {
	cfg := config.GetConfig()

	useridInt, strconvErr := strconv.Atoi(userid)
	if strconvErr != nil {
		return domain.Interests{}, strconvErr
	}
	filter := bson.M{"userid": useridInt}
	result := d.DB.Database(cfg.DB_NAME).Collection("interests").FindOne(ctx, filter)

	// Convert string slice to objectID slice
	var newObjIdSlice []primitive.ObjectID
	for _, val := range req.Tags {
		objid, objconvErr := primitive.ObjectIDFromHex(val)
		if objconvErr != nil {
			return domain.Interests{}, errors.New("invalid tag id found")
		}

		newObjIdSlice = append(newObjIdSlice, objid)
	}

	// If no document found, create new document
	if result.Err() == mongo.ErrNoDocuments {
		useridInt, convErr := strconv.Atoi(userid)
		if convErr != nil {
			return domain.Interests{}, convErr
		}
		newUserIntrest := domain.Interests{
			ID:     primitive.NewObjectID(),
			UserID: uint(useridInt),
			Tags:   newObjIdSlice,
		}
		_, err := d.DB.Database(cfg.DB_NAME).Collection("interests").InsertOne(ctx, newUserIntrest)
		if err != nil {
			return domain.Interests{}, err
		}

		return newUserIntrest, nil
	}

	// If document found, update the document
	var updatedIntrest domain.Interests
	opts := options.FindOneAndUpdate().SetUpsert(false).SetReturnDocument(options.After)
	res := d.DB.Database(cfg.DB_NAME).Collection("interests").FindOneAndUpdate(ctx, filter, bson.M{"$push": bson.M{"tags": bson.M{"$each": newObjIdSlice}}}, opts)

	if res.Err() != nil {
		return domain.Interests{}, fmt.Errorf("update failed: %w", res.Err())
	}

	if err := res.Decode(&updatedIntrest); err != nil {
		return domain.Interests{}, fmt.Errorf("result decoding failed: %w", err)
	}

	// Return updated document
	return updatedIntrest, nil
}

func (d *ContentUserDatabase) FetchRecommendedPosts(ctx context.Context, userid int64) ([]domain.Contents, error) {
	cfg := config.GetConfig()

	// // Fetch user interests
	// intrestFilter := bson.M{"userid": userid}
	// intrestCursor, intrestFindErr := d.DB.Database(cfg.DB_NAME).Collection("interests").Find(ctx, intrestFilter)
	// if intrestFindErr != nil {
	// 	return []domain.Contents{}, fmt.Errorf("find error: %w", intrestFindErr)
	// }
	// defer intrestCursor.Close(ctx)
	// var intrests []domain.Interests

	// for intrestCursor.Next(ctx) {
	// 	var intrest domain.Interests
	// 	if err := intrestCursor.Decode(&intrest); err != nil {
	// 		return []domain.Contents{}, fmt.Errorf("decode error: %w", err)
	// 	}
	// 	intrests = append(intrests, intrest)
	// }

	// for i := 0; i < len(intrests); i++ {
	// 	fmt.Printf("%+v\n", intrests[i])
	// }

	// if len(intrests) <= 0 {
	// 	return []domain.Contents{}, errors.New("follow at least one tag")
	// }

	// var recomentPosts []domain.Contents

	// collection := d.DB.Database(cfg.DB_NAME).Collection("contents")

	// for i := 0; i < len(intrests); i++ {
	// 	fmt.Printf("%+v", intrests[i])
	// 	filter := bson.M{"labels": intrests[i]}
	// 	cursor, err := collection.Find(ctx, filter)
	// 	if err != nil {
	// 		return []domain.Contents{}, err
	// 	}

	// 	// Iterate over the cursor to fetch all matching documents
	// 	for cursor.Next(ctx) {
	// 		var result domain.Contents
	// 		if err := cursor.Decode(&result); err != nil {
	// 			return []domain.Contents{}, err
	// 		}

	// 		recomentPosts = append(recomentPosts, result)
	// 	}
	// 	cursor.Close(ctx)

	// }

	var userIntrests []domain.Interests
	filter := bson.M{"userid": userid}
	cursor, err := d.DB.Database(cfg.DB_NAME).Collection("interests").Find(ctx, filter)

	if err != nil {
		return []domain.Contents{}, err
	}

	// Iterate over the cursor to fetch all matching documents
	for cursor.Next(ctx) {
		var result domain.Interests
		if err := cursor.Decode(&result); err != nil {
			return []domain.Contents{}, err
		}

		userIntrests = append(userIntrests, result)
	}
	cursor.Close(ctx)

	//
	tagFilter := bson.M{"_id": bson.M{"$in": userIntrests[0].Tags}}
	tagCursor, tagerr := d.DB.Database(cfg.DB_NAME).Collection("tags").Find(ctx, tagFilter)
	if tagerr != nil {
		return []domain.Contents{}, tagerr
	}
	defer tagCursor.Close(ctx)

	var recomentedContentIds []primitive.ObjectID
	for tagCursor.Next(ctx) {
		var currentTag domain.Tags
		if err := tagCursor.Decode(&currentTag); err != nil {
			return []domain.Contents{}, err
		}

		recomentedContentIds = append(recomentedContentIds, currentTag.Posts...)
	}

	//
	var recomentedDocs []domain.Contents
	postFilter := bson.M{"_id": bson.M{"$in": recomentedContentIds}}
	postCuresor, postFinderr := d.DB.Database(cfg.DB_NAME).Collection("contents").Find(ctx, postFilter)
	if postCuresor.Err() != nil {
		return []domain.Contents{}, postFinderr
	}

	for postCuresor.Next(ctx) {
		var post domain.Contents
		if err := postCuresor.Decode(&post); err != nil {
			return []domain.Contents{}, err
		}

		recomentedDocs = append(recomentedDocs, post)
	}

	if err := postCuresor.Err(); err != nil {
		return []domain.Contents{}, err
	}

	defer postCuresor.Close(ctx)

	return recomentedDocs, nil
}

func (d *ContentUserDatabase) FetchAllTags(ctx context.Context) ([]domain.Tags, error) {
	cfg := config.GetConfig()
	collection := d.DB.Database(cfg.DB_NAME).Collection("tags")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []domain.Tags{}, err
	}
	defer cursor.Close(ctx)

	var tags []domain.Tags
	if err = cursor.All(ctx, &tags); err != nil {
		return []domain.Tags{}, err
	}

	return tags, nil
}

func (d *ContentUserDatabase) GetOnePost(ctx context.Context, postid string) (domain.Contents, error) {
	cfg := config.GetConfig()
	objid, objErr := primitive.ObjectIDFromHex(postid)
	if objErr != nil {
		return domain.Contents{}, objErr
	}
	filter := bson.M{"_id": objid}

	// fetch user data on user service
	resp, userErr := http.NewRequest("GET", "http://muarif.online/api/v1/users/account", nil)
	if userErr != nil {
		return domain.Contents{}, userErr
	}
	var userVal requests.UserValue
	client := http.Client{}
	user, user_err := client.Do(resp)
	if user_err != nil {
		return domain.Contents{}, user_err
	}

	userValbyte, readErr := io.ReadAll(user.Body)
	if readErr != nil {
		return domain.Contents{}, readErr
	}

	if err := json.Unmarshal(userValbyte, &userVal); err != nil {
		return domain.Contents{}, err
	}

	var post domain.Contents
	err := d.DB.Database(cfg.DB_NAME).Collection("contents").FindOne(ctx, filter).Decode(&post)
	if err != nil {
		return domain.Contents{}, err
	}

	if post.IsPremium && !userVal.Data.Is_premium {
		n := len(post.Body)
		post.Body = post.Body[:n/4] + "......... || " + "this is premium content please buy premium to read full content"
	}

	return post, nil
}
