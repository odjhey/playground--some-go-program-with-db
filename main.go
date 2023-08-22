package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/halakata/go-pokemon-api/http_api"
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

	r.Method("GET", "/x", Handler(http_api.SomeHandler))
	r.Method("GET", "/message", Handler(http_api.GetMessage))
	r.Method("POST", "/message", Handler(http_api.CreateMessage))
	return r

}
