package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// func query_handler() {

// 	query_string := "SELECT \"ResId\", \"ResName\" FROM tbl_dk_resource"
// 	query_string = "SELECT * FROM tbl_dk_res_category"
// 	do_db_query(query_string)

// }

type QueryRequest struct {
	QueryString string `json:"query_string"`
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
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

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func withParams(w http.ResponseWriter, r *http.Request) {
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

func handleRequests() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", notFound)
	api.HandleFunc("/user/{userID}/comment/{commentID}", withParams).Methods(http.MethodGet)

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	handleRequests()
}
