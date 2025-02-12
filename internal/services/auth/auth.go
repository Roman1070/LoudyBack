package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"loudy-back/internal/domain/models"
	"loudy-back/internal/lib/jwt"
	"loudy-back/internal/lib/logger/sl"
	storage "loudy-back/internal/storage"
	"loudy-back/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	log          *slog.Logger
	userSaver    UserSaver
	userProvider UserProvider
	tokenTTL     time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passwordHash []byte) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidAppID       = errors.New("invalid app id")
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
)

// New returns a new instance of the Auth service.
func New(log *slog.Logger, userProvider UserProvider, userSaver UserSaver, tokenTTL time.Duration) *AuthService {
	return &AuthService{
		userSaver:    userSaver,
		userProvider: userProvider,
		log:          log,
		tokenTTL:     tokenTTL,
	}
}

func (a *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := a.userProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrAppNotFound) {
			a.log.Warn("user not found", sl.Err(err))

			return "", fmt.Errorf("servic Login error: " + ErrInvalidCredentials.Error())
		}

		a.log.Error("failed to get user", sl.Err(err))
		return "", fmt.Errorf("service Login error: " + ErrInvalidCredentials.Error())
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		a.log.Info("invalid credentials", sl.Err(err))

		return "", fmt.Errorf("servic Login error: " + ErrInvalidCredentials.Error())
	}

	token, err := jwt.NewToken(user, a.tokenTTL)
	if err != nil {
		a.log.Error("Failed to generate token", sl.Err(err))
		return "", fmt.Errorf("servic Login error: " + err.Error())
	}

	return token, nil
}

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
