package storage

import "errors"

var (
	ErrUserExists          = errors.New("user already exists")
	ErrUserNotFound        = errors.New("user not found")
	ErrArtistNotFound      = errors.New("artist not found")
	ErrProfileNotFound     = errors.New("profile not found")
	ErrTrackNotFound       = errors.New("track not found")
	ErrArtistAlreadyExists = errors.New("artist already exists")
	ErrAlbumNotFound       = errors.New("album not found")
	ErrAlbumAlreadyExists  = errors.New("album already exists")
)
