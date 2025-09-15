package repos

import (
	"database/sql"
)

type AuthRepo interface {
	GetPassword(string) (string, error)
}

type AuthRepoImpl struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepo {
	return &AuthRepoImpl{
		db: db,
	}
}

func (ar *AuthRepoImpl) GetPassword(email string) (string, error) {
	var passHash string

	results, err := ar.db.Query("SELECT password FROM users where email = ?", email)
	if err != nil {
		return "", err
	}
	defer results.Close()

	results.Next()
	if err := results.Scan(&passHash); err != nil {
		return "", err
	}

	return passHash, nil
}
