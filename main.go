package main

import (
	"log"
	"os"
	"strconv"

	"concepts/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	name := os.Args[1]
	rollnum := os.Args[2]
	branch := os.Args[3]

	rollno, err := strconv.Atoi(rollnum)
	if err != nil {
		log.Fatal("Error converting rollnum to integer: ", err)
	}

	//create database
	db, err := database.CreateDb()
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	//create students table
	err = database.CreateTable(db)
	if err != nil{
		log.Fatal("Error creating table students: ", err)
	}

	err = database.StoreInDb(db, name, branch, rollno)
	if err != nil {
		log.Fatal("Error inserting data to table students: ", err)
	}

	err = database.FetchDetails(db, "Hellen")
	if err != nil {
		log.Fatal("Error fetching details from database: ", err)
	}
}
