package models

type CreateAlbumRequest struct {
	Name        string   `json:"name"`
	Cover       string   `json:"cover"`
	ReleaseDate string   `json:"release_date"`
	ArtistsIds  []string `json:"artists_ids"`
}
