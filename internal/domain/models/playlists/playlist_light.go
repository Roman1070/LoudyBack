package models

import (
	models "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlaylistLight struct {
	ID          primitive.ObjectID
	Name        string
	CreatorID   primitive.ObjectID
	CreatorName string
	Tracks      []models.TrackLight
}
