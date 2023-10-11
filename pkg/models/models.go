package models

import (
	"time"
)

// Models are the representation of the entities in the database.
// If you add a new one, remember to add it to the Purge and Migrate functions in pkg/repository/database.go

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	IsAdmin   bool
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
