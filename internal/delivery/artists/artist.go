package artists

import (
	"encoding/json"
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"
	"loudy-back/utils"
	"net/http"
)

func (c *ArtistsClient) Artist(w http.ResponseWriter, r *http.Request) {
	slog.Info("client start [Artist]")
	name := r.URL.Query().Get("name")

	artist, err := c.ArtistsGRPCClient.Artist(r.Context(), &artistsv1.ArtistRequest{
		Name: name,
	})
	if err != nil {
		slog.Error("[Artist] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	result, err := json.Marshal(artist)
	if err != nil {
		slog.Error("[Artist] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
