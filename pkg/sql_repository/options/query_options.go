package options

import (
	"gorm.io/gorm"
)

// QueryOptions are used to modify the DB object & query string
type QueryOption func(db *gorm.DB, query *string) *gorm.DB

// Filters
func WithoutDeleted() QueryOption {
	return func(db *gorm.DB, query *string) *gorm.DB {
		*query += " AND deleted = false"
		return db
	}
}

func WithUsername(username string) QueryOption {
	return func(db *gorm.DB, query *string) *gorm.DB {
		if username == "" {
			return db
		}
		return db.Where("username LIKE ?", "%"+username+"%")
	}
}

// Preloaders
func WithDetails() QueryOption {
	return func(db *gorm.DB, query *string) *gorm.DB {
		return db.Preload("Details")
	}
}

func WithPosts() QueryOption {
	return func(db *gorm.DB, query *string) *gorm.DB {
		return db.Preload("Posts")
	}
}
