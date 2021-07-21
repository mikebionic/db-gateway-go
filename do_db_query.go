package main

import (
	"database/sql"
	"fmt"
	"log"
)

func do_db_query(query_string string) {
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

	query_request := query_string

	rows, err := db.Query(query_request)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cols, _ := rows.Columns()

	var query_results []interface{}

	for rows.Next() {

		values := make([]interface{}, len(cols))
		valuePointers := make([]interface{}, len(cols))

		for i := range cols {
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			log.Fatal(err)
		}

		for i, col := range cols {
			query_result_dict := map[interface{}]interface{}{}

			val := values[i]
			b, ok := val.([]byte)

			var v interface{}

			if ok {
				v = string(b)
			} else {
				v = val
			}
			query_result_dict[col] = v

			fmt.Println(query_result_dict)
			query_results = append(query_results, query_result_dict)
		}

		fmt.Println("resutls---------")
		fmt.Println(query_results)

	}

	defer db.Close()
	// return nil
}
