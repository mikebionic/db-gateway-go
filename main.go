package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// func query_handler() {

// 	query_string := "SELECT \"ResId\", \"ResName\" FROM tbl_dk_resource"
// 	query_string = "SELECT * FROM tbl_dk_res_category"
// 	do_db_query(query_string)

// }

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", r))
}
