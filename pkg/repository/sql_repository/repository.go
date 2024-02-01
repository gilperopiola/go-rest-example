package sql_repository

type repository struct {
	*Database
}

func New(database *Database) *repository {
	return &repository{Database: database}
}
