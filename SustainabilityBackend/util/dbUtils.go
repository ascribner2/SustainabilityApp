package util

import "database/sql"

func GetDbResultsSlice[T any](results *sql.Rows, startingCapacity int64) []T {
	rows := make([]T, 0, startingCapacity)

	for i := 0; results.Next(); i++ {
		var row T
		results.Scan(&row)
		rows[int64(i)] = row
	}

	return rows
}
