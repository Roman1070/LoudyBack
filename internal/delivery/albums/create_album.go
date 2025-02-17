package albums

import (
	"encoding/json"
	albumsv1 "loudy-back/gen/go/albums"
	models "loudy-back/internal/domain/models/albums"
	"loudy-back/utils"
	"net/http"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c *AlbumsClient) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	c.log.Info("[CreateAlbum] client started")

	var request models.CreateAlbumRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		c.log.Error("[CreateAlbum] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	_, err = c.ALbumsGRPCClient.CreateAlbum(r.Context(), &albumsv1.CreateAlbumRequest{
		Name:        request.Name,
		ArtistsIds:  request.ArtistsIds,
		Cover:       request.Cover,
		ReleaseDate: timestamppb.New(request.ReleaseDate),
	})

	if err != nil {
		c.log.Error("[CreateAlbum] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
	}

	w.WriteHeader(http.StatusOK)
}
