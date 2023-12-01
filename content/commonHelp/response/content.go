package response

import (
	"time"

	"github.com/muhammedarifp/content/domain"
)

type ContentVal struct {
	CreateAt        time.Time         `bson:"create_at"`
	LastUpdate      time.Time         `bson:"last_update"`
	UserID          string            `bson:"user_id"`
	ThumbnailImg    string            `bson:"thumbnail_img"`
	Title           string            `bson:"title"`
	Body            string            `bson:"body"`
	Like            uint              `bson:"like"`
	IsShowReactions bool              `bson:"is_show_reactions"`
	IsPremium       bool              `bson:"is_premium"`
	Comments        []domain.Comment  `bson:"comments"`
	Reactions       []domain.Reaction `bson:"reactions"`
}
