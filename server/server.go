package main

import (
	"flag"
	"log"
	"os"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var databasePath string
var rebuildDatabase bool

func init() {
	flag.StringVar(&databasePath, "database", "db.sqlite3", "Database file path")
	flag.BoolVar(&rebuildDatabase, "rebuild", false, "Rebuild database")
}

func main() {
	log.Println("Server start")
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal("Open database failed: ", err)
	}
	initDatabase(db)
	StartServer()
}

func initDatabase(db *sql.DB) {
	if rebuildDatabase {
		os.Remove(databasePath)
	}
	_, err := os.Stat(databasePath)
	if err == nil || os.IsExist(err) {
		log.Printf("Reading database from file %s ...\n", databasePath)
		return
	}
	// init database
	log.Println("Building database ...")
	initStmt, err := db.Prepare(
		`CREATE TABLE 'suggestion' (
		'id' INTEGER,
		'type' INTEGER,
		'time' INTEGER,
		'content' TEXT
	)`)
	if err != nil {
		log.Fatal("Error at create suggestion table: ", err)
	}
	_, err = initStmt.Exec()
	if err != nil {
		log.Fatal("Error at create suggestion table: ", err)
	}
}
