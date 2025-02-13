package main

import (
	"encoding/json"
	"log/slog"
	contentv1 "loudy-back/gen/go/content"
	models "loudy-back/internal/domain/models/content"
	"loudy-back/utils"
	"net/http"
)

type ContentClient struct {
	contentProvider ContentProvider
	contentCreator  contentv1.ContentClient
}

type ContentProvider interface {
	Artist(name string) (models.Artist, error)
	Album(id uint32) (models.Album, error)
	SearchContent(input string) ([]models.ArtistLight, []models.AlbumLight, []models.TrackLight, error)
}

func (c *ContentClient) Artist(w http.ResponseWriter, r *http.Request) {
	slog.Info("client start [Artist]")
	name := r.URL.Query().Get("name")

	artist, err := c.contentProvider.Artist(name)
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

func (c *ContentClient) CreateArtist(w http.ResponseWriter, r *http.Request) {
	slog.Info("client start [CreateArtist]")

	var request models.CreateArtistRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.Error("[CreateArtist] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	resp, err := c.contentCreator.CreateArtist(r.Context(), request.ToGRPC())
	if err != nil {
		slog.Error("[CreateArtist] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	result, err := json.Marshal(resp)
	if err != nil {
		slog.Error("[CreateArtist] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
