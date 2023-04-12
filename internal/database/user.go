package database

import (
	"github.com/go-faster/errors"
	"golang.org/x/crypto/bcrypt"
)

func (db Database) UserChangePassword(userID uint32, oldPassword, newPassword string) error {
	// Get the old user's password
	user := &User{ID: userID}
	err := db.db.First(user).Error
	if err != nil {
		return errors.Wrap(err, "cannot get user")
	}
	// Check the password
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(oldPassword))
	if err != nil {
		return ErrInvalidCredentials
	}
	// Update the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "cannot hash password")
	}
	user.HashedPassword = string(hashedPassword)
	err = db.db.Save(user).Error
	if err != nil {
		return errors.Wrap(err, "cannot update password")
	}
	return nil
}

// GetUser will get all info about a user
func (db Database) GetUser(userID uint32) (User, error) {
	user := User{ID: userID}
	err := db.db.First(&user).Error
	if err != nil {
		return User{}, errors.Wrap(err, "cannot query user")
	}
	user.HashedPassword = "" // we remove this from json. But just to be sure
	return user, nil
}

// UpdateAbout will update the about section of a user
func (db Database) UpdateAbout(userID uint32, about string) error {
	user := User{ID: userID}
	err := db.db.Model(&user).Update("about", about).Error
	if err != nil {
		return errors.Wrap(err, "cannot update about of user")
	}
	return nil
}
