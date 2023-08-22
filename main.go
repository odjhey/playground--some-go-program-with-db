package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type SomeMessage struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func (s *SomeMessage) Bind(r *http.Request) error {
	if s.Message == "" {
		return errors.New("message is required")
	}

	return nil
}

func (s *SomeMessage) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type ErrResponse struct {
	Err            error  `json:"-"`
	HttpStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	AppCode        int64  `json:"code,omitempty"`
	ErrorText      string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HttpStatusCode)
	return nil
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here
		w.WriteHeader(500)
		w.Write([]byte("invalid"))
	}
}

func main() {

	r := getRouter()
	http.ListenAndServe(":8080", r)

}

func getRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	r.Method("GET", "/x", Handler(someHandler))
	r.Method("GET", "/message", Handler(getMessage))
	r.Method("POST", "/message", Handler(createMessage))
	return r

}

func createMessage(w http.ResponseWriter, r *http.Request) error {

	message := &SomeMessage{}

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

type FileContentMessage struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

type FileContentMessageData struct {
	Data []FileContentMessage `json:"data"`
}

func renderBadRequest(message string) *ErrResponse {
	return &ErrResponse{
		HttpStatusCode: 400,
		StatusText:     message,
	}
}

func getMessage(w http.ResponseWriter, r *http.Request) error {

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
		render.Render(w, r, renderBadRequest("Invalid id"))
		return nil
	}

	id, err := strconv.Atoi(idQuery)
	if err != nil || id < 0 {
		render.Render(w, r, renderBadRequest("Invalid id"))
		return nil
	}

	for i := range messages.Data {
		if messages.Data[i].ID == id {
			render.Render(w, r, &SomeMessage{
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

func someHandler(w http.ResponseWriter, r *http.Request) error {
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
