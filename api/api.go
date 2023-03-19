package api

import (
	"RayaneshBackend/internal/database"
	"RayaneshBackend/pkg/session"
)

type API struct {
	Database database.Database
	Session  session.Session
}
