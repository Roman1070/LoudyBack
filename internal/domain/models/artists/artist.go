package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Artist struct {
	Name       string
	Albums     []primitive.ObjectID
	Cover      string
	Bio        string
	LikesCount uint32 `bson:"likes_count"`
}
