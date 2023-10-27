package options

import (
	"github.com/jinzhu/gorm"
)

// QueryOptions are used to modify the query string
type QueryOption func(*string)

func WithoutDeleted(query *string) {
	*query += " AND deleted = false"
}

// PreloadOptions are used to preload fields on a search query
type PreloadOption func(*gorm.DB) *gorm.DB

func WithDetails(db *gorm.DB) *gorm.DB {
	return db.Preload("Details")
}

func WithPosts(db *gorm.DB) *gorm.DB {
	return db.Preload("Posts")
}
