package database

import (
	"github.com/go-faster/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

// NewDatabase will create a new database connection backed by gorm
func NewDatabase(dsn string) (Database, error) {
	// Open the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return Database{}, errors.Wrap(err, "cannot connect to database")
	}
	// Get the connection to ping it
	genericDB, err := db.DB()
	if err != nil {
		return Database{}, errors.Wrap(err, "cannot get generic db object")
	}
	err = genericDB.Ping()
	if err != nil {
		return Database{}, errors.Wrap(err, "cannot ping database")
	}
	// Migrate
	err = db.AutoMigrate(&User{})
	if err != nil {
		return Database{}, errors.Wrap(err, "cannot migrate database")
	}
	// Done
	return Database{db}, nil
}

// Close will close the database
func (db Database) Close() {
	genericDB, _ := db.db.DB()
	if genericDB != nil {
		_ = genericDB.Close()
	}
}
