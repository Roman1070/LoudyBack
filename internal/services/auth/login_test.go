package auth

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"loudy-back/internal/domain/models"
	mock_auth "loudy-back/internal/services/auth/mocks"
	storage "loudy-back/internal/storage"
	"loudy-back/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthService_Login(t *testing.T) {
	type want struct {
		token string
		err   error
	}

	var tests = []struct {
		name      string
		email     string
		password  string
		setupFunc func(ctrl *gomock.Controller) *AuthService
		want      want
	}{
		{
			name:     "Успешная авторизация",
			email:    "test@mail.ru",
			password: "password",

			setupFunc: func(ctrl *gomock.Controller) *AuthService {
				userSaver := mock_auth.NewMockUserSaver(ctrl)
				userProvider := mock_auth.NewMockUserProvider(ctrl)
				salt, _ := utils.GenerateSalt()
				user := models.User{
					ID:           1,
					Email:        "test@mail.ru",
					PasswordHash: utils.HashPassword("password", []byte(salt)),
				}

				userProvider.EXPECT().User(gomock.Any(), "test@mail.ru").Return(user, nil)

				return &AuthService{
					userSaver:    userSaver,
					userProvider: userProvider,
					log:          slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				token: "some_token",
				err:   nil,
			},
		}, {
			name:     "Неверный пароль",
			email:    "test@mail.ru",
			password: "password",

			setupFunc: func(ctrl *gomock.Controller) *AuthService {
				userSaver := mock_auth.NewMockUserSaver(ctrl)
				userProvider := mock_auth.NewMockUserProvider(ctrl)
				salt, _ := utils.GenerateSalt()
				user := models.User{
					ID:           1,
					Email:        "test@mail.ru",
					PasswordHash: utils.HashPassword("other_password", []byte(salt)),
				}

				userProvider.EXPECT().User(gomock.Any(), "test@mail.ru").Return(user, nil)

				return &AuthService{
					userSaver:    userSaver,
					userProvider: userProvider,
					log:          slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				token: "",
				err:   errors.New("service Login error: " + ErrInvalidCredentials.Error()),
			},
		}, {
			name:     "Пользователь не найден",
			email:    "test@mail.ru",
			password: "password",

			setupFunc: func(ctrl *gomock.Controller) *AuthService {
				userSaver := mock_auth.NewMockUserSaver(ctrl)
				userProvider := mock_auth.NewMockUserProvider(ctrl)

				userProvider.EXPECT().User(gomock.Any(), "test@mail.ru").Return(models.User{}, storage.ErrUserNotFound)

				return &AuthService{
					userSaver:    userSaver,
					userProvider: userProvider,
					log:          slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				token: "",
				err:   errors.New("service Login error: " + ErrUserNotFound.Error()),
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

			token, err := service.Login(ctx, tt.email, tt.password)

			assert.NotNil(t, tt.want.token, token)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
