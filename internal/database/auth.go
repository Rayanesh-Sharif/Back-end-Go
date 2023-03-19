package database

import (
	"github.com/go-faster/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

// UserLogin will try to log in a user and check for their email and password.
// This will return the user object.
func (db Database) UserLogin(email, password string) (User, error) {
	// Query user
	user := User{Email: email}
	err := db.db.First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, ErrInvalidCredentials
		}
		return User{}, errors.Wrap(err, "cannot query user")
	}
	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return User{}, ErrInvalidCredentials
	}
	return user, nil
}
