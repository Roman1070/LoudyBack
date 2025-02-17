package models

import (
	"time"
)

type CreateAlbumRequest struct {
	Name        string    `json:"name"`
	Cover       string    `json:"cover"`
	ReleaseDate time.Time `json:"release_date"`
	ArtistsIds  []string  `json:"artists_ids"`
}
