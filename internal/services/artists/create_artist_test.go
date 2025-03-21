package artists

import (
	"context"
	"errors"
	"io"
	"log/slog"
	models "loudy-back/internal/domain/models/artists"
	"loudy-back/internal/storage"
	"testing"

	mock_artists "loudy-back/internal/services/artists/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/emptypb"
)

type key int

const (
	testContextRequestIDKey key = iota
)

var (
	testContextRequestIDValue = uuid.New()
)

func TestArtistservice_CreateArtist(t *testing.T) {
	type want struct {
		err error
	}

	var tests = []struct {
		name       string
		artistName string
		cover      string
		bio        string
		albums     []primitive.ObjectID
		likesCount uint32
		setupFunc  func(ctrl *gomock.Controller) *ArtistsService
		want       want
	}{
		{
			name:       "Успешное создание артиста ",
			artistName: "artist name",
			cover:      "",
			bio:        "",
			setupFunc: func(ctrl *gomock.Controller) *ArtistsService {
				artistsService := mock_artists.NewMockArtists(ctrl)

				artistsService.EXPECT().Artist(gomock.Any(), "artist name").Return(models.Artist{}, storage.ErrArtistNotFound)

				artistsService.EXPECT().CreateArtist(gomock.Any(), "artist name", "", "").Return(&emptypb.Empty{}, nil)

				return &ArtistsService{
					artists: artistsService,
					log:     slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				err: nil,
			},
		},
		{
			name:       "Артист уже существует",
			artistName: "artist name",
			cover:      "",
			bio:        "",
			setupFunc: func(ctrl *gomock.Controller) *ArtistsService {
				artistsService := mock_artists.NewMockArtists(ctrl)

				artist := models.Artist{
					Name:       "artist name",
					Bio:        "",
					Cover:      "",
					LikesCount: 0,
				}

				artistsService.EXPECT().Artist(gomock.Any(), "artist name").Return(artist, nil)

				artistsService.EXPECT().CreateArtist(gomock.Any(), "artist name", "", "").
					Return(&emptypb.Empty{}, storage.ErrArtistAlreadyExists).AnyTimes()

				return &ArtistsService{
					artists: artistsService,
					log:     slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				err: storage.ErrArtistAlreadyExists,
			},
		},
		{
			name:       "Внутренняя ошибка",
			artistName: "artist name",
			cover:      "",
			bio:        "",
			setupFunc: func(ctrl *gomock.Controller) *ArtistsService {
				sertistsService := mock_artists.NewMockArtists(ctrl)

				sertistsService.EXPECT().Artist(gomock.Any(), "artist name").Return(models.Artist{}, storage.ErrArtistNotFound)

				sertistsService.EXPECT().CreateArtist(gomock.Any(), "artist name", "", "").
					Return(&emptypb.Empty{}, errors.New("some error"))

				return &ArtistsService{
					artists: sertistsService,
					log:     slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				err: errors.New("[CreateArtist] service error: some error"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.WithValue(context.Background(), testContextRequestIDKey, testContextRequestIDValue)

			service := tt.setupFunc(ctrl)

			_, err := service.CreateArtist(ctx, tt.artistName, tt.cover, tt.bio)

			assert.Equal(t, tt.want.err, err)
		})
	}
}
