package main

import (
	_ "github.com/lib/pq"
)

func main() {

	query_string := "SELECT \"ResId\", \"ResName\" FROM tbl_dk_resource"
	query_string = "SELECT * FROM tbl_dk_res_category"
	do_db_query(query_string)

	// fmt.Println(response)

}
