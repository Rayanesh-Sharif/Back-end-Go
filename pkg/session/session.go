package session

import (
	"crypto/rand"
	"encoding/base32"
)

type Session struct {
	Storage
}

func NewSession(storage Storage) Session {
	return Session{storage}
}

func newToken() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return base32.StdEncoding.EncodeToString(b)
}
