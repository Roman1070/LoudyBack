package artists

import (
	"encoding/json"
	"fmt"
	"log/slog"
	models "loudy-back/internal/domain/models/artists"
	"loudy-back/internal/storage"
	"loudy-back/utils"
	"net/http"
	"strings"
)

func (c *ArtistsClient) CreateArtist(w http.ResponseWriter, r *http.Request) {
	slog.Info("[CreateArtist] client started ")

	var request models.CreateArtistRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.Error("[CreateArtist] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	_, err = c.ArtistsGRPCClient.CreateArtist(r.Context(), request.ToGRPC())
	if err != nil {
		slog.Error("[CreateArtist] client error: " + err.Error())
		if strings.Contains(err.Error(), storage.ErrArtistAlreadyExists.Error()) {
			utils.WriteError(w, fmt.Sprintf("Artist %v already exists.", request.Name))
		} else {
			utils.WriteError(w, "Internal error")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
