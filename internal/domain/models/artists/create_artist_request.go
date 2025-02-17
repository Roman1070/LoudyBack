package models

import (
	artistsv1 "loudy-back/gen/go/artists"
)

type CreateArtistRequest struct {
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Cover string `json:"cover"`
}

func (r *CreateArtistRequest) ToGRPC() *artistsv1.CreateArtistRequest {
	return &artistsv1.CreateArtistRequest{
		Name:  r.Name,
		Bio:   r.Bio,
		Cover: r.Cover,
	}
}
