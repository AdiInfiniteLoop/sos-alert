package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"sos-alert/db"
	"sos-alert/internal/handlers"
)

const (
	HTTPStatusOK = http.StatusOK
)

func main() {
	fmt.Println("Welcome to SOS Alert")
	db.OpenDatabaseConnection() //Opening database connection
	defer db.CloseDatabaseConnection()
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(HTTPStatusOK)
		_, err := w.Write([]byte("Welcome to SOS-Alert"))
		if err != nil {
			log.Fatal("Cannot get route", err)
		}
	})

	v1Router := chi.NewRouter()
	router.Mount("/api/v1", v1Router)
	v1Router.Post("/", handlers.PostAlert) //Expects a function reference not the function itself

	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error Occured while serving", err)
	}
}
