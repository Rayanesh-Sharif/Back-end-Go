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

// UserRegister will register the user in database
func (db Database) UserRegister(email, username, password string) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "cannot hash the password")
	}
	// Save the password in database
	result := db.db.Create(&User{
		Username:       username,
		HashedPassword: string(hashedPassword),
		Email:          email,
	})
	if result.Error != nil {
		return errors.Wrap(result.Error, "cannot insert into database")
	}
	// Done
	return nil
}
