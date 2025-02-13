package models

import contentv1 "loudy-back/gen/go/content"

type CreateArtistRequest struct {
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Cover string `json:"cover"`
}

func (r *CreateArtistRequest) ToGRPC() *contentv1.CreateArtistRequest {
	return &contentv1.CreateArtistRequest{
		Name:  r.Name,
		Bio:   r.Bio,
		Cover: r.Cover,
	}
}
