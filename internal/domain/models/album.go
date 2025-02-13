package models

import "time"

type Album struct {
	ID          uint32
	Artists     []ArtistLight
	Name        string
	Cover       string
	Tracks      []TrackLight
	ReleaseDate time.Time
}
