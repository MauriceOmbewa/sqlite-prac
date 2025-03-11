package database

import "database/sql"

func StoreInDb(db *sql.DB, name, branch string, rollno int) error {
	_, err := db.Exec(`INSERT INTO students(name, rollno, branch) VALUES(?, ?, ?)`, name, rollno, branch)
	if err != nil {
		return err
	}
	return nil
}