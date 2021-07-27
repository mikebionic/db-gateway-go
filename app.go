package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(db_type, db_user, db_password, db_host, db_database string) {

	connStr := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", db_type, db_user, db_password, db_host, db_database)

	var err error
	a.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	fmt.Printf("Server running on  %s \n", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	api := a.Router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", a.getRequest).Methods(http.MethodGet)
	api.HandleFunc("/make-db-request", a.apiMakeDbRequest).Methods(http.MethodPost)
	api.HandleFunc("/", a.notFound)
	api.HandleFunc("/user/{userID}/comment/{commentID}", a.withParams).Methods(http.MethodGet)
}
