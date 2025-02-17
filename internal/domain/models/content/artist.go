package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Artist struct {
	ID         primitive.ObjectID
	Name       string
	Albums     []AlbumLight
	Cover      string
	Bio        string
	LikesCount uint32
}
