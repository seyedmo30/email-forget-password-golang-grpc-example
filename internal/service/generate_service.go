package service

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func GenerateCode() uint32 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Uint32()
}

func AddGenerateCode(email *string, code *uint32) error {

	db, _ := sql.Open("sqlite3", "../../database.db")
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")

	return r1.Uint32()
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"program" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")
}
#TODO ابتدا توی اسکیولایت جدول بساز با توجه به فانک بالا سپس انسرت کن با لینک پایین
https://www.codeproject.com/Articles/5261771/Golang-SQLite-Simple-Example
