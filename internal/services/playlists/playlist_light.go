package playlists

import (
	"context"
	models "loudy-back/internal/domain/models/playlists"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *PlaylistsService) PlaylistLight(ctx context.Context, id primitive.ObjectID) (models.PlaylistLight, error) {
	return s.playlists.PlaylistLight(ctx, id)
}
