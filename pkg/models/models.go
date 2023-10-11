package models

import (
	"time"
)

// Models are the representation of the entities in the database.
// If you add a new one, remember to add it to the Purge and Migrate functions in pkg/repository/database.go

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	IsAdmin   bool
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
