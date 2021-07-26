package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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
	api.HandleFunc("", a.getRequest).Methods(http.MethodGet)
	api.HandleFunc("", a.apiMakeDbRequest).Methods(http.MethodPost)
	api.HandleFunc("", a.notFound)
	api.HandleFunc("/user/{userID}/comment/{commentID}", a.withParams).Methods(http.MethodGet)

}

type QueryRequest struct {
	QueryString string `json:"query_string"`
}

func (a *App) apiMakeDbRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)

	var queryrequest QueryRequest
	json.Unmarshal(reqBody, &queryrequest)
	fmt.Println(queryrequest)

	response, err := do_db_query(queryrequest.QueryString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error"}`))
	}

	fmt.Println(response)
	type RespObject struct {
		Data []interface{}
	}
	respObject := RespObject{response}
	data, err := json.Marshal(respObject)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("", data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	// var res map[string]interface{}
	// json.Marshal(&response, &res)

	// w.Write([]byte(`{"message": "post called"}`))
}

func (a *App) getRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func (a *App) withParams(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	commentID := -1
	if val, ok := pathParams["commentID"]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	query := r.URL.Query()
	location := query.Get("location")
	// example query
	// http://127.0.0.1:8080/api/v1/user/23/comment/55?location=elsewhere
	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}

func (a *App) notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
