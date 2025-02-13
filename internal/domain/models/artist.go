package models

type Artist struct {
	ID         uint32
	Albums     []AlbumLight
	Cover      string
	Bio        string
	LikesCount uint32
}
