package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here
		w.WriteHeader(500)
		w.Write([]byte("invalid"))
	}
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	r.Method("GET", "/x", Handler(someHandler))

	http.ListenAndServe(":8080", r)

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
