package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type QueryRequest struct {
	QueryString   string   `json:"query_string"`
	Base64Columns []string `json:"base64_columns"`
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

	var response []query_response

	switch executeOnly_state {
	case 0:
		response, err = do_db_select_query(a.DB, queryrequest.QueryString, queryrequest.Base64Columns)
	default:
		err = do_db_query_exec(a.DB, queryrequest.QueryString)
	}
	fmt.Print(response, err)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := make(map[string]interface{})
	res["message"] = "db query result"
	res["data"] = response

	if executeOnly_state == 0 {
		res["total"] = len(response)
	} else {
		res["total"] = 1
		res["status"] = 1
	}

	respondWithJSON(w, http.StatusOK, res)
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
