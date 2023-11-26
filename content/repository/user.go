package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/muhammedarifp/content/commonHelp/response"
	"github.com/muhammedarifp/content/config"
	"github.com/muhammedarifp/content/domain"
	"github.com/muhammedarifp/content/repository/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContentUserDatabase struct {
	DB *mongo.Client
}

func NewContentUserRepo(db *mongo.Client) interfaces.ContentUserRepository {
	return &ContentUserDatabase{DB: db}
}

func (d *ContentUserDatabase) CreateNewPost(ctx context.Context) (response.ContentResp, error) {
	cfg := config.GetConfig()
	var emptyContentResp response.ContentResp

	time.Sleep(time.Second * 5)

	select {
	case <-ctx.Done():
		return emptyContentResp, errors.New("time limit reached")
	default:
		time.Sleep(time.Second)
	}

	tempContent := `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>Hello</h1>
    <img src="https://techexchangeblog.s3.ap-south-1.amazonaws.com/profile/1.jpg" alt="">
    <p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Maiores in quaerat ex recusandae similique odio praesentium ratione non error dolor reprehenderit harum ab, obcaecati aliquam doloribus quis modi accusantium! Fugiat, dolorem! Culpa iure, neque adipisci perspiciatis dolores esse placeat, non aperiam, nisi exercitationem odio saepe illum inventore explicabo pariatur ipsa?</p>
</body>
</html>
	`
	newContent := domain.Contents{
		CreateAt:        time.Now(),
		LastUpdate:      time.Now(),
		UserID:          "1",
		ThumbnailImg:    "",
		Title:           "Sample",
		Body:            tempContent,
		Like:            0,
		IsShowReactions: true,
		IsPremium:       false,
		Comments:        []domain.Comment{},
		Reactions:       []domain.Reaction{},
	}
	res, err := d.DB.Database(cfg.DB_NAME).Collection("contents").InsertOne(context.TODO(), newContent)
	if err != nil {
		return emptyContentResp, err
	}

	fmt.Println(res)

	return emptyContentResp, nil
}
