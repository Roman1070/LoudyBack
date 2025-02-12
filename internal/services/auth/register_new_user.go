package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"loudy-back/internal/lib/logger/sl"
	storage "loudy-back/internal/storage"
	"loudy-back/utils"
)

func (a *AuthService) RegisterNewUser(ctx context.Context, email string, password string) (int64, error) {
	slog.Info("service started to RegisterNewUser")

	salt, err := utils.GenerateSalt()
	if err != nil {
		a.log.Error("Failed to generate salt", sl.Err(err))
		return -1, fmt.Errorf("service RegisterNewUser error: " + err.Error())
	}

	hashedPassword := utils.HashPassword(password, []byte(salt))

	id, err := a.userSaver.SaveUser(ctx, email, hashedPassword)
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			slog.Warn("User already exists")

			return -1, fmt.Errorf("service RegisterNewUser error: " + storage.ErrUserExists.Error())
		}
		slog.Error("Failed to save user", sl.Err(err))
		return -1, fmt.Errorf("service RegisterNewUser error: " + err.Error())
	}

	return id, nil
}
