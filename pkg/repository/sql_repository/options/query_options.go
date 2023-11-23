package options

import (
	"gorm.io/gorm"
)

/*---------------------------------------------------------------------
// QueryOptions are used to modify the DB object or the query string.
// Variadic functions approach.
//--------------------------------*/

type QueryOption func(db *gorm.DB, query *string) *gorm.DB

func ApplyQueryOptions(db *gorm.DB, query *string, opts ...any) *gorm.DB {
	for _, opt := range opts {
		if optFn, ok := opt.(QueryOption); ok {
			db = optFn(db, query)
		}
	}
	return db
}

/*--------------------------
//     Query Filters
//------------------------*/

func WithoutDeleted() QueryOption {
	return func(db *gorm.DB, query *string) *gorm.DB {
		*query += " AND deleted = false"
		return db
	}
}

func WithUsername(username string) QueryOption {
	return func(db *gorm.DB, query *string) *gorm.DB {
		if username != "" {
			return db.Where("username LIKE ?", "%"+username+"%")
		}
		return db
	}
}

/*--------------------------
//    Query Preloaders
//------------------------*/

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
