package database

// User contains the information about a user
type User struct {
	ID             uint32
	Username       string `gorm:"not null"`
	HashedPassword string `gorm:"not null"`
	Email          string `gorm:"uniqueIndex;not null"`
}
