package db

import (
	"database/sql"
	"errors"

	"github.com/halakata/go-pokemon-api/models"
)

func (db Database) GetAllMessages() ([]models.SomeMessage, error) {
	list := []models.SomeMessage{}

	rows, err := db.Conn.Query(`SELECT * FROM "ta1" LIMIT 50;`)
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var message models.SomeMessage
		err := rows.Scan(&message.ID, &message.Message)
		if err != nil {
			return list, err
		}
		list = append(list, message)
	}
	return list, nil

}

func (db Database) GetMessageById(messageId int) (models.SomeMessage, error) {
	message := models.SomeMessage{}

	query := `SELECT * FROM "ta1" WHERE id = $1;`
	row := db.Conn.QueryRow(query, messageId)

	switch err := row.Scan(&message.ID, &message.Message); err {
	case sql.ErrNoRows:
		return message, errors.New("No Match")
	default:
		return message, err

	}

}

func (db Database) CreateMessage(inputId int, inputMessage string) error {

	var id int
	var message string
	query := `INSERT INTO "ta1" (id, message) VALUES ($1, $2) RETURNING id, message`
	err := db.Conn.QueryRow(query, inputId, inputMessage).Scan(&id, &message)
	if err != nil {
		return err
	}
	return nil

}
