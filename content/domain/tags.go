package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tags struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"	id"`
	CreateAt time.Time            `bson:"create_at" json:"create_at"`
	UpdateAt time.Time            `bson:"update_at" json:"update_at"`
	Tag      string               `bson:"tag" json:"tag"`
	Posts    []primitive.ObjectID `bson:"posts" json:"posts"`
	IsActive bool                 `bson:"is_active"`
}

type Interests struct {
	ID     primitive.ObjectID   `bson:"_id,omitempty" json:"	id"`
	UserID uint                 `bson:"userid,omitempty" json:"userid"`
	Tags   []primitive.ObjectID `bson:"tags" json:"tags"`
}
