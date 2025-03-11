package database

import (
	"database/sql"
	"fmt"
)

type Student struct{
	ID int
	Name string
	Rollno int
	Branch string
}

func FetchDetails(db *sql.DB, name string) error{
	var student Student
	err := db.QueryRow(`SELECT id, name, rollno, branch FROM students WHERE name=?`, name).
	Scan(&student.ID, &student.Name, &student.Rollno, &student.Branch)
	if err != nil{
		return err
	}

	fmt.Printf("ID: %v, Name: %v, RollNo: %v, Branch: %v\n", student.ID, student.Name, student.Rollno, student.Branch)
	return nil
}