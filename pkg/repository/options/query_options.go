package options

import "github.com/jinzhu/gorm"

type QueryOption func(*string)

func WithoutDeleted(query *string) {
	*query += " AND deleted = false"
}

type PreloadOption func(*gorm.DB) *gorm.DB

func WithDetails(db *gorm.DB) *gorm.DB {
	return db.Preload("Details")
}

func WithPosts(db *gorm.DB) *gorm.DB {
	return db.Preload("Posts")
}
