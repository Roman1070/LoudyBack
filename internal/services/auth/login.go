package auth

import (
	"context"
	"errors"
	"fmt"
	"loudy-back/internal/lib/jwt"
	"loudy-back/internal/lib/logger/sl"
	storage "loudy-back/internal/storage"
	"loudy-back/utils"
)

func (a *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := a.userProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", sl.Err(err))

			return "", fmt.Errorf("%s", "service Login error: "+ErrUserNotFound.Error())
		}

		a.log.Error("failed to get user", sl.Err(err))
		return "", fmt.Errorf("%s", "service Login error: "+ErrInvalidCredentials.Error())
	}

	if !utils.VerifyPassword(string(user.PasswordHash), password) {
		a.log.Info("invalid credentials")
		return "", fmt.Errorf("%s", "service Login error: "+ErrInvalidCredentials.Error())
	}

	token, err := jwt.NewToken(user, a.tokenTTL)
	if err != nil {
		a.log.Error("Failed to generate token", sl.Err(err))
		return "", fmt.Errorf("%s", "service Login error: "+err.Error())
	}

	return token, nil
}
