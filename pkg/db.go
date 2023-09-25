package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_NAME = "notes.sdb"
)

type TDb struct {
	db *sql.DB
}

// Create database in file 
func (t *TDb) Init(filename string) {
	db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	t.db = db
	t.createSchema()
}

// create base schema for notes database 
func (t *TDb) createSchema() {
	sql := "CREATE TABLE IF NOT EXISTS tb_notes(id unsigned integer primary key desc, note text)" //id is the date of the note
	_, error := t.db.Exec(sql)
	if error != nil {
		log.Print("Error to create table")
	}
	return
}

// Save a note
func (t *TDb) SaveNote(note string) bool {
	id := time.Now().Unix()
	sql := fmt.Sprintf("INSERT INTO tb_notes(id, note) VALUES (%d,'%s')", id, note)
	_, error := t.db.Exec(sql)
	if error != nil {
		log.Print(error)
		return false
	}
	return true
}

// search note by text, use fulltext search
func (t *TDb) SearchNote(text string) (*sql.Rows, bool) {
	sql := fmt.Sprintf("SELECT * FROM tb_notes WHERE note like '%%%s%%'", text)
	rows, err := t.db.Query(sql)
	if err != nil {
		log.Print("No rows")
		return rows, false
	}
	return rows, true
}

// TDb constructor
func NewDb() *TDb {
	db := new(TDb)
	db.Init(DB_NAME)
	return db
}
