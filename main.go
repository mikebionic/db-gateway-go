package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	db_type := "postgres"
	db_user := "postgres"
	db_password := "123456"
	db_host := "localhost"
	db_database := "dbSapHasap"

	connStr := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", db_type, db_user, db_password, db_host, db_database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	query_request := "SELECT \"ResId\", \"ResName\" FROM tbl_dk_resource"

	rows, err := db.Query(query_request)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	cols, _ := rows.Columns()

	for rows.Next() {

		values := make([]interface{}, len(cols))
		valuePointers := make([]interface{}, len(cols))

		for i := range cols {
			valuePointers[i] = &values[i]
		}

		err := rows.Scan(valuePointers...)
		if err != nil {
			log.Fatal(err)
		}

		for i, col := range cols {

			val := values[i]

			b, ok := val.([]byte)
			var v interface{}

			if ok {
				v = string(b)
			} else {
				v = val
			}

			fmt.Println(col, v)

		}

	}

	defer db.Close()

}
