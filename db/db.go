package db

import (
	"database/sql"
	"fmt"
	"log"
)

const DbContextKey = "dbContext"

type Database struct {
	Conn *sql.DB
}

func Init(username, password, database, host string, port int) (Database, error) {

	db := Database{}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)

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
