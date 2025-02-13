package postgre

import (
	"context"
	"fmt"
	"log/slog"
	models "loudy-back/internal/domain/models/content"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Storage) Artist(ctx context.Context, name string) (models.Artist, error) {
	slog.Info("storage start [Artist]")

	const query = `
		SELECT id,name,bio,likes_count,cover
		FROM artists 
		WHERE name = $1; 
	`
	const albumsIdsQuery = `
		SELECT album_id 
		FROM artists_albums 
		WHERE artist_id = $1;
	`

	var artist models.Artist
	err := s.db.QueryRow(ctx, query, name).Scan(&artist.ID, &artist.Name, &artist.Bio, &artist.LikesCount, &artist.Cover)
	if err != nil {
		slog.Error("storage [Artist] error: " + err.Error())
		return models.Artist{}, fmt.Errorf("storage [Artist] error: " + err.Error())
	}

	rows, err := s.db.Query(ctx, albumsIdsQuery, artist.ID)
	if err != nil {
		slog.Error("storage [Artist] error: " + err.Error())
		return models.Artist{}, fmt.Errorf("storage [Artist] error: " + err.Error())
	}

	defer rows.Close()
	albums_ids := make([]any, 0, 12)
	for rows.Next() {
		var id uint32
		err = rows.Scan(&id)
		if err != nil {
			slog.Error("storage [Artist] error: " + err.Error())
			return models.Artist{}, fmt.Errorf("storage [Artist] error: " + err.Error())
		}

		albums_ids = append(albums_ids, id)
	}

	artist.Albums, err = s.GetArtistsAlbumsLight(ctx, albums_ids)
	if err != nil {
		slog.Error("storage [Artist] error: " + err.Error())
		return models.Artist{}, fmt.Errorf("storage [Artist] error: " + err.Error())
	}

	return artist, nil
}

func (s *Storage) GetArtistsAlbumsLight(ctx context.Context, ids []any) ([]models.AlbumLight, error) {
	slog.Info("storage start [GetArtistsAlbumsLight]")

	idsRequestString := "("
	i := 1
	for ; i < len(ids); i++ {
		idsRequestString += fmt.Sprintf("$%v,", i)
	}

	idsRequestString += fmt.Sprintf("$%v)", i)
	query := fmt.Sprintf(`
		SELECT id,name,cover,release_date
		FROM albums 
		WHERE id IN %v;`, idsRequestString)

	rows, err := s.db.Query(ctx, query, ids...)
	if err != nil {
		slog.Error("storage [GetArtistsAlbumsLight] error: " + err.Error())
		return nil, fmt.Errorf("storage [GetArtistsAlbumsLight] error: " + err.Error())
	}

	albums := make([]models.AlbumLight, 0, 12)
	for rows.Next() {
		var album models.AlbumLight
		var date time.Time
		err = rows.Scan(&album.ID, &album.Name, &album.Cover, &date)
		if err != nil {
			slog.Error("storage [GetArtistsAlbumsLight] error: " + err.Error())
			return nil, fmt.Errorf("storage [GetArtistsAlbumsLight] error: " + err.Error())
		}

		album.Year = uint32(date.Year())
		albums = append(albums, album)
	}
	return albums, nil
}

func (s *Storage) Album(ctx context.Context, id uint32) (models.Album, error) {
	return models.Album{}, nil
}
func (s *Storage) SearchContent(ctx context.Context, input string) ([]models.ArtistLight, []models.AlbumLight, []models.TrackLight, error) {
	return nil, nil, nil, nil
}
func (s *Storage) CreateArtist(ctx context.Context, name, cover, bio string) (uint32, error) {
	return 0, nil
}

func (s *Storage) CreateAlbum(ctx context.Context, name, cover string, tracks_ids []uint32, releaseDate time.Time) (uint32, error) {
	return 0, nil
}

func (s *Storage) CreateTrack(ctx context.Context, name, file string, albumId uint32) (*emptypb.Empty, error) {
	return nil, nil
}
