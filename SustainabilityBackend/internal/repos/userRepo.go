package repos

import "database/sql"

type UserRepo interface {
	InsertUser(string, string) error
}

type UserRepoImpl struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}

func (ur *UserRepoImpl) InsertUser(email string, password string) error {
	_, err := ur.db.Query("INSERT INTO Users (email, password) VALUES (?, ?)", email, password)
	if err != nil {
		return err
	}
	return nil
}
