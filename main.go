package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/halakata/go-pokemon-api/db"
	"github.com/halakata/go-pokemon-api/http_api"

	_ "github.com/lib/pq"
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

	addr := ":8080"

	dbUser, dbPassword, dbName, dbHost, dbPort :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT")

	dbPortInt, _ := strconv.Atoi(dbPort)

	database, err := db.Init(dbUser, dbPassword, dbName, dbHost, dbPortInt)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Conn.Close()

	dbInstance = database
	r := getRouter()
	http.ListenAndServe(addr, r)

}

var dbInstance db.Database

func DbContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), db.DbContextKey, dbInstance)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(DbContext)
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	r.Method("GET", "/x", Handler(http_api.SomeHandler))
	r.Method("GET", "/message", Handler(http_api.GetMessage))
	r.Method("POST", "/message", Handler(http_api.CreateMessage))
	r.Method("GET", "/message-db", Handler(http_api.GetMessageFromDb))
	r.Method("POST", "/message-db", Handler(http_api.CreateMessageDb))
	return r

}

// DO WE NEED TO ADD `func Stop` for graceful?
