package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contents struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	CreateAt        time.Time          `bson:"create_at"`
	LastUpdate      time.Time          `bson:"last_update"`
	UserID          uint               `bson:"user_id"`
	ThumbnailImg    string             `bson:"thumbnail_img"`
	Title           string             `bson:"title"`
	Body            string             `bson:"body"`
	Like            uint               `bson:"like"`
	IsShowReactions bool               `bson:"is_show_reactions"`
	IsPremium       bool               `bson:"is_premium"`
	Comments        []Comment          `bson:"comments"`
	Reactions       []Reaction         `bson:"reactions"`
}

type LabelMapping struct {
	CreateAt time.Time  `bson:""`
	Label    string     `bson:""`
	Contents []Contents `bson:""`
}

// Helpers

type Comment struct {
	CreateAt time.Time `bson:"create_at"`
	UserID   string    `bson:"user_id"`
	Message  string    `bson:"message"`
}

type Reaction struct {
	CreateAt time.Time `bson:"create_at"`
	UserID   string    `bson:"user_id"`
}
