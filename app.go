package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(db_type, db_user, db_password, db_host, db_database, db_params string) {

	var connStr string
	switch db_type {
	case "sqlserver":
		connStr = fmt.Sprintf("%s://%s:%s@%s/instance?database=%s&%s", db_type, db_user, db_password, db_host, db_database, db_params)
	default:
		connStr = fmt.Sprintf("%s://%s:%s@%s/%s?%s", db_type, db_user, db_password, db_host, db_database, db_params)
	}
	fmt.Print(connStr)

	var err error
	a.DB, err = sql.Open(db_type, connStr)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	fmt.Printf("Server running on  %s \n", addr)
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	// ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(addr, handlers.CORS(credentials, methods, origins)(a.Router)))
}

func (a *App) initializeRoutes() {
	api := a.Router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", a.getRequest).Methods(http.MethodGet)
	api.HandleFunc("/make-db-request", a.apiMakeDbRequest).Methods(http.MethodPost)
	api.HandleFunc("/", a.notFound)
	api.HandleFunc("/user/{userID}/comment/{commentID}", a.withParams).Methods(http.MethodGet)
}
