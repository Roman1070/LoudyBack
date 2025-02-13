package models

type Artist struct {
	ID         uint32
	Name       string
	Albums     []AlbumLight
	Cover      string
	Bio        string
	LikesCount uint32
}
