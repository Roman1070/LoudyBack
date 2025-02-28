package tracks

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *TracksStorage) CreateTrack(ctx context.Context, name, filename string, albumId primitive.ObjectID, artistsIds []primitive.ObjectID, duration uint16) (primitive.ObjectID, error) {
	s.log.Info("[CreateTrack] storage started")

	track := dtoTrack{
		ID:         primitive.NewObjectID(),
		Name:       name,
		Filename:   filename,
		AlbumId:    albumId,
		ArtistsIds: artistsIds,
		Duration:   duration,
	}

	_, err := s.collection.InsertOne(ctx, track)
	if err != nil {
		s.log.Error("[CreateTrack] storage error: " + err.Error())
		return [12]byte{}, errors.New("[CreateTrack] storage error: " + err.Error())
	}

	return track.ID, nil
}
