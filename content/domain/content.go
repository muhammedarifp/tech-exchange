package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contents struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreateAt        time.Time          `bson:"create_at" json:"create_at"`
	LastUpdate      time.Time          `bson:"last_update" json:"last_update"`
	UserID          uint               `bson:"user_id" json:"user_id"`
	ThumbnailImg    string             `bson:"thumbnail_img" json:"thumbnail_img"`
	Title           string             `bson:"title" json:"title"`
	Body            string             `bson:"body" json:"body"`
	Like            uint               `bson:"like" json:"like"`
	IsShowReactions bool               `bson:"is_show_reactions"  json:"is_show_reactions"`
	IsPremium       bool               `bson:"is_premium" json:"is_premium"`
	Comments        []Comment          `bson:"comments" json:"comments"`
	Reactions       []Reaction         `json:"reactions" bson:"reactions"`
	IsActive        bool               `json:"is_active" bson:"is_active"`
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
