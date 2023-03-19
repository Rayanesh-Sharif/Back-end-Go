package session

import (
	"github.com/go-faster/errors"
	"time"
)

var ErrNotFound = errors.New("not found")

type Storage interface {
	Store(userID uint32, ttl time.Duration) (accessToken, refreshToken string, err error)
	Get(accessToken string) (userID uint32, err error)
	Delete(accessToken string) error
	Refresh(refreshToken string, ttl time.Duration) (newAccessToken string, err error)
}
