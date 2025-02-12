package content

import (
	"log/slog"
	"time"
)

type ContentService struct {
	log             *slog.Logger
	contentProvider ContentProvider
	tokenTTL        time.Duration
}

type ContentProvider interface {
}

// New returns a new instance of the Auth service.
func New(log *slog.Logger, contentProvider ContentProvider, tokenTTL time.Duration) *ContentService {
	return &ContentService{
		contentProvider: contentProvider,
		log:             log,
		tokenTTL:        tokenTTL,
	}
}
