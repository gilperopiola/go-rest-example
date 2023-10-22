package options

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// QueryOptions are used to modify the query string
type QueryOption func(*string)

func WithoutDeleted(query *string) {
	*query += " AND deleted = false"
}

// PreloadOptions are used to preload fields on a search query
type PreloadOption func(DB) DB

func WithDetails(db DB) DB {
	return db.Preload("Details")
}

func WithPosts(db DB) DB {
	return db.Preload("Posts")
}

type DB interface {
	Create(value interface{}) *gorm.DB
	Preload(column string, conditions ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Model(value interface{}) *gorm.DB
	Update(attrs ...interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	Offset(offset interface{}) *gorm.DB
	Limit(limit interface{}) *gorm.DB
	Close() error
	LogMode(enable bool) *gorm.DB
	DB() *sql.DB
	AutoMigrate(values ...interface{}) *gorm.DB
	DropTable(values ...interface{}) *gorm.DB
}
