package db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	HOST = "localhost"
	PORT = 5432
)

const DbContextKey = "dbContext"

type Database struct {
	Conn *sql.DB
}

func Init(username, password, database string) (Database, error) {

	db := Database{}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)

	conn, err := sql.Open("postgres", dsn)

	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	log.Println("Database connection established")
	return db, nil
}
