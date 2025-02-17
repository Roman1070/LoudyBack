package albums

import (
	"encoding/json"
	albumsv1 "loudy-back/gen/go/albums"
	"loudy-back/utils"
	"net/http"
)

func (c *AlbumsClient) Album(w http.ResponseWriter, r *http.Request) {
	c.log.Info("[Album] client started")

	resp, err := c.ALbumsGRPCClient.Album(r.Context(), &albumsv1.AlbumRequest{
		Id: r.URL.Query().Get("id"),
	})
	if err != nil {
		c.log.Error("[Album] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	result, err := json.Marshal(resp)
	if err != nil {
		c.log.Error("[Album] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
