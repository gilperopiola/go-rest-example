package options

import (
	"go.mongodb.org/mongo-driver/bson"
)

type FilterOption func(filter bson.M) bson.M

// GetFilterFromOptions returns a filter from the given FilterOptions.
func GetFilterFromOptions(opts ...any) bson.M {
	filter := bson.M{}
	for _, opt := range opts {
		if optFn, ok := opt.(FilterOption); ok {
			filter = optFn(filter)
		}
	}
	return filter
}

// WithoutDeleted adds a filter to exclude deleted documents

// WithUsername adds a filter for username
func WithUsername(username string) FilterOption {
	return func(filter bson.M) bson.M {
		if username != "" {
			filter["username"] = username
		}
		return filter
	}
}
