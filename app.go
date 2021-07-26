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

////////

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

////////

type QueryRequest struct {
	QueryString string `json:"query_string"`
}

func (a *App) apiMakeDbRequest(w http.ResponseWriter, r *http.Request) {

	executeOnly := r.URL.Query().Get("executeOnly")
	executeOnly_state := 0
	var err error

	if len(executeOnly) > 0 {
		executeOnly_state, err = strconv.Atoi(executeOnly)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	var queryrequest QueryRequest
	json.Unmarshal(reqBody, &queryrequest)
	fmt.Println(queryrequest)

	var response interface{}
	switch executeOnly_state {
	case 0:
		response, err = do_db_select_query(a.DB, queryrequest.QueryString)

	default:
		err = do_db_query_exec(a.DB, queryrequest.QueryString)
		response = map[string]interface{}{"status": 1}

	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, response)

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
