package storage

import "errors"

var (
	ErrUserExists          = errors.New("user already exists")
	ErrUserNotFound        = errors.New("user not found")
	ErrArtistNotFound      = errors.New("artist not found")
	ErrArtistAlreadyExists = errors.New("artist already exists")
)
