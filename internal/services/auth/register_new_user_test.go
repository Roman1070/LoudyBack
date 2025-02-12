package auth

import (
	"context"
	"errors"
	"io"
	"log/slog"
	mock_auth "loudy-back/internal/services/auth/mocks"
	"loudy-back/internal/storage"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type key int

const (
	testContextRequestIDKey key = iota
)

var (
	testContextRequestIDValue = uuid.New()
)

func TestAuthService_Register(t *testing.T) {
	type want struct {
		id  int64
		err error
	}

	var tests = []struct {
		name      string
		email     string
		password  string
		setupFunc func(ctrl *gomock.Controller) *AuthService
		want      want
	}{
		{
			name:     "Успешное создание пользователя",
			email:    "test@mail.ru",
			password: "password",

			setupFunc: func(ctrl *gomock.Controller) *AuthService {
				userSaver := mock_auth.NewMockUserSaver(ctrl)
				userProvider := mock_auth.NewMockUserProvider(ctrl)

				returnedId := int64(1)
				userSaver.EXPECT().SaveUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(returnedId, nil)

				return &AuthService{
					userSaver:    userSaver,
					userProvider: userProvider,
					log:          slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				id:  1,
				err: nil,
			},
		}, {
			name:     "Ошибка: пользователь уже есть",
			email:    "test@mail.ru",
			password: "password",

			setupFunc: func(ctrl *gomock.Controller) *AuthService {
				userSaver := mock_auth.NewMockUserSaver(ctrl)
				userProvider := mock_auth.NewMockUserProvider(ctrl)

				userSaver.EXPECT().SaveUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(-1), storage.ErrUserExists)

				return &AuthService{
					userSaver:    userSaver,
					userProvider: userProvider,
					log:          slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				id:  -1,
				err: errors.New("service RegisterNewUser error: " + storage.ErrUserExists.Error()),
			},
		}, {
			name:     "Ошибка при создании пользователя",
			email:    "test@mail.ru",
			password: "password",

			setupFunc: func(ctrl *gomock.Controller) *AuthService {
				userSaver := mock_auth.NewMockUserSaver(ctrl)
				userProvider := mock_auth.NewMockUserProvider(ctrl)

				userSaver.EXPECT().SaveUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(-1), errors.New("some error"))

				return &AuthService{
					userSaver:    userSaver,
					userProvider: userProvider,
					log:          slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				id:  -1,
				err: errors.New("service RegisterNewUser error: some error"),
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

			user, err := service.RegisterNewUser(ctx, tt.email, tt.password)

			assert.NotNil(t, tt.want.id, user)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
