package artists

import (
	models "loudy-back/internal/domain/models/artists"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dtoArtist struct {
	ID         primitive.ObjectID   `bson:"_id"`
	Name       string               `bson:"name"`
	Cover      string               `bson:"cover"`
	Bio        string               `bson:"bio"`
	LikesCount uint32               `bson:"likes_count"`
	AlbumsIds  []primitive.ObjectID `bson:"albums_ids"`
}

func (artist *dtoArtist) toCommonModel() models.Artist {
	return models.Artist{
		ID:         artist.ID,
		Name:       artist.Name,
		Cover:      artist.Cover,
		Bio:        artist.Bio,
		LikesCount: uint32(artist.LikesCount),
		Albums:     artist.AlbumsIds,
	}
}
