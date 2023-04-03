package api

import (
	"RayaneshBackend/internal/database"
	"RayaneshBackend/pkg/session"
)

type API struct {
	Database database.Database
	Session  session.Session
	// StaticFolder is the folder which we store our static files in it
	StaticFolder string
}
