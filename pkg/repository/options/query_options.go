package options

type QueryOption func(*string)

func WithoutDeleted(query *string) {
	*query += " AND deleted = false"
}
