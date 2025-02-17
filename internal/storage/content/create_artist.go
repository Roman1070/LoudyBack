package content

import (
	"context"
	"fmt"
	"log/slog"
	models "loudy-back/internal/domain/models/content"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ContentStorage) CreateArtist(ctx context.Context, name, cover, bio string) (*emptypb.Empty, error) {
	newArtist := models.Artist{
		Name:       name,
		Cover:      cover,
		Bio:        bio,
		LikesCount: 0,
		Albums:     []models.AlbumLight{},
	}

	result, err := s.collection.InsertOne(ctx, newArtist)
	if err != nil {
		slog.Error("[CreateArtist] storage error: " + err.Error())
		return nil, fmt.Errorf("%s", "[CreateArtist] storage error: "+err.Error())
	}

	s.log.Info("[CreateArtist] storafe finished successfully, id=" + fmt.Sprint(result.InsertedID))
	return nil, nil
}
