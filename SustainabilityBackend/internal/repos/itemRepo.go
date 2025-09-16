package repos

import (
	"database/sql"

	"github.com/ascribner/sustainabilityapp/internal/entity"
)

type ItemRepo interface {
	InsertItem(string, float64, string, string) error
	GetItems(string) ([]entity.Item, error)
}

type ItemRepoImpl struct {
	db *sql.DB
}

func NewItemRepo(db *sql.DB) ItemRepo {
	return &ItemRepoImpl{
		db: db,
	}
}

func (ir *ItemRepoImpl) InsertItem(title string, offset float64, date string, user_email string) error {
	_, err := ir.db.Query("INSERT INTO items (title, offset, date, user_email) VALUES (?, ?, ?, ?)", title, offset, date, user_email)
	if err != nil {
		return err
	}

	return nil
}

func (ir *ItemRepoImpl) GetItems(user_email string) ([]entity.Item, error) {
	items := make([]entity.Item, 0, 10)

	results, err := ir.db.Query("SELECT title, offset, date FROM items WHERE user_email = ?", user_email)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var title string
		var offset float64
		var date string

		results.Scan(&title, &offset, &date)

		row := entity.Item{
			Title:  title,
			Offset: offset,
			Date:   date,
		}
		items = append(items, row)
	}

	return items, nil
}
