package playlists

import (
	"context"
	"errors"
	models "loudy-back/internal/domain/models/playlists"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *PlaylistsService) Playlist(ctx context.Context, id primitive.ObjectID) (models.Playlist, error) {
	s.log.Info("[Playlist] service started")

	preliminary, err := s.playlists.Playlist(ctx, id)
	if err != nil {
		s.log.Error("[Playlist] service error: " + err.Error())
		return models.Playlist{}, errors.New("[Playlist] service error: " + err.Error())
	}

	tracks, err := s.tracksProvider.Tracks(ctx, preliminary.TracksIds)
	if err != nil {
		s.log.Error("[Playlist] service error: " + err.Error())
		return models.Playlist{}, errors.New("[Playlist] service error: " + err.Error())
	}

	return models.Playlist{
		ID:          id,
		Name:        preliminary.Name,
		Cover:       preliminary.Cover,
		CreatorID:   preliminary.CreatorID,
		CreatorName: preliminary.CreatorName,
		Tracks:      tracks,
	}, nil
}
