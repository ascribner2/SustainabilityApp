package repos

import "database/sql"

type AuthRepo interface {
	getPassword(string) (string, error)
}

type AuthRepoImpl struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepo {
	return &AuthRepoImpl{
		db: db,
	}
}

func (ar *AuthRepoImpl) getPassword(email string) (string, error) {
	results, err := ar.db.Query("SELECT password FROM users where email = ?", email)
	if err != nil {
		return "", err
	}

	rows := map[int64]string{}

	for i := 0; results.Next(); i++ {
		var row string
		results.Scan(&row)
		rows[int64(i)] = row
	}

	return rows[0], nil
}
