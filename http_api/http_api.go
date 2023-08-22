package http_api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/render"
	"github.com/halakata/go-pokemon-api/db"
	"github.com/halakata/go-pokemon-api/models"
)

type FileContentMessage struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

type FileContentMessageData struct {
	Data []FileContentMessage `json:"data"`
}

func GetMessage(w http.ResponseWriter, r *http.Request) error {

	content, err := os.ReadFile("./static/messages.json")
	if err != nil {
		render.Render(w, r, &ErrResponse{
			Err:            err,
			HttpStatusCode: 404,
			StatusText:     "Not found",
			ErrorText:      err.Error(),
		})
		return nil
	}

	var messages FileContentMessageData
	json.Unmarshal(content, &messages)

	idQuery := r.URL.Query().Get("id")
	if idQuery == "" {
		render.Render(w, r, RenderBadRequest("Invalid id"))
		return nil
	}

	id, err := strconv.Atoi(idQuery)
	if err != nil || id < 0 {
		render.Render(w, r, RenderBadRequest("Invalid id"))
		return nil
	}

	for i := range messages.Data {
		if messages.Data[i].ID == id {
			render.Render(w, r, &models.SomeMessage{
				ID:      messages.Data[i].ID,
				Message: messages.Data[i].Message,
			})
			return nil
		}
	}

	render.Render(w, r, &ErrResponse{
		HttpStatusCode: 404,
		StatusText:     "Not found",
	})

	return nil
}

func CreateMessage(w http.ResponseWriter, r *http.Request) error {

	message := &models.SomeMessage{}

	if err := render.Bind(r, message); err != nil {
		render.Render(w, r, &ErrResponse{
			Err:            err,
			HttpStatusCode: 500,
			StatusText:     "something wrong haha",
			ErrorText:      err.Error()})
		return nil
	}

	render.Render(w, r, message)

	return nil

}

func SomeHandler(w http.ResponseWriter, r *http.Request) error {
	idQuery := r.URL.Query().Get("id")
	if idQuery == "" {
		return errors.New(idQuery)
	}

	id, err := strconv.Atoi(idQuery)
	if err != nil || id < 0 {
		return errors.New(idQuery)
	}

	w.Write([]byte("some name " + strconv.Itoa(id)))
	return nil
}

func GetMessageFromDb(w http.ResponseWriter, r *http.Request) error {

	dbInstance := r.Context().Value(db.DbContextKey).(db.Database)

	idQuery := r.URL.Query().Get("id")
	if idQuery == "" {
		render.Render(w, r, RenderBadRequest("Invalid id"))
		return nil
	}

	id, err := strconv.Atoi(idQuery)
	if err != nil || id < 0 {
		render.Render(w, r, RenderBadRequest("Invalid id"))
		return nil
	}

	row, err := dbInstance.GetMessageById(id)
	if err != nil {
		render.Render(w, r, &ErrResponse{
			HttpStatusCode: 404,
			StatusText:     "Not found",
			ErrorText:      err.Error(),
		})
	}

	render.Render(w, r, &models.SomeMessage{
		ID:      row.ID,
		Message: row.Message,
	})

	return nil
}
