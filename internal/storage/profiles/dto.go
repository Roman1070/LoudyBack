package profiles

import "go.mongodb.org/mongo-driver/bson/primitive"

type dtoProfile struct {
	ID                primitive.ObjectID   `bson:"_id"`
	UserId            uint32               `bson:"user_id"`
	Name              string               `bson:"name"`
	Avatar            string               `bson:"avatar"`
	Bio               string               `bson:"bio"`
	LikesCount        uint32               `bson:"likes_count"`
	SavedTracksIds    []primitive.ObjectID `bson:"saved_tracks_ids"`
	SavedAlbumsIds    []primitive.ObjectID `bson:"saved_albums_ids"`
	SavedArtistsIds   []primitive.ObjectID `bson:"saved_artists_ids"`
	SavedPlaylistsIds []primitive.ObjectID `bson:"saved_playlists_ids"`
}
