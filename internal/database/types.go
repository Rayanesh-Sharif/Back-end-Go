package database

// User contains the information about a user
type User struct {
	ID             uint32
	Username       string
	HashedPassword string
	Email          string `gorm:"uniqueIndex"`
}
