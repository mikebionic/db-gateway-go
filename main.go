package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// func query_handler() {

// 	query_string := "SELECT \"ResId\", \"ResName\" FROM tbl_dk_resource"
// 	query_string = "SELECT * FROM tbl_dk_res_category"
// 	do_db_query(query_string)

// }

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
