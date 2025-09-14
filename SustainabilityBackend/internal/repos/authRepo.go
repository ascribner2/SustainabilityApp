package repos

import (
	"database/sql"

	"github.com/ascribner/sustainabilityapp/util"
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
	results, err := ar.db.Query("SELECT password FROM users where email = ?", email)
	if err != nil {
		return "", err
	}
	defer results.Close()

	rows := util.GetDbResultsSlice[string](results, 1)

	return rows[0], nil
}
