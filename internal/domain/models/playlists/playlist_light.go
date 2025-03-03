package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlaylistLight struct {
	ID          primitive.ObjectID
	Name        string
	CreatorID   primitive.ObjectID
	CreatorName string
	Cover       string
}
