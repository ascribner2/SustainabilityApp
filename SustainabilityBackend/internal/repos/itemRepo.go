package repos

import "database/sql"

type ItemRepo interface {
}

type ItemRepoImpl struct {
	db *sql.DB
}

func NewItemRepo(db *sql.DB) ItemRepo {
	return &ItemRepoImpl{
		db: db,
	}
}
