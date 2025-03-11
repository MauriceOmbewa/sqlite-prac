package database

import "database/sql"

func CreateDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "student.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}