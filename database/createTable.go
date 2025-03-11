package database

import "database/sql"

func CreateTable(db *sql.DB) error{
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS students(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT[50] NOT NULL, rollno INTEGER UNIQUE NOT NULL, branch TEXT[50] NOT NULL)`)
	if err != nil {
		return err
	}
	return nil
}