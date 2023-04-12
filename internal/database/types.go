package database

// User contains the information about a user
type User struct {
	ID             uint32 `json:"id"`
	Username       string `gorm:"not null" json:"username"`
	HashedPassword string `gorm:"not null" json:"-"`
	Email          string `gorm:"uniqueIndex;not null" json:"email"`
	About          string `gorm:"not null;default:''" json:"about"`
}
