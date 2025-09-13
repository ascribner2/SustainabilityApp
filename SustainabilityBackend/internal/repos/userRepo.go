package repos

import "database/sql"

type UserRepo interface {
}

type UserRepoImpl struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}
