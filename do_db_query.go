package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type query_response map[string]interface{}

func do_db_query(db *sql.DB, query_string string) ([]query_response, error) {

	rows, err := db.Query(query_string)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	query_results := []query_response{}

	for rows.Next() {

		r := make(query_response)
		cols, _ := rows.Columns()

		fmt.Println(cols, "\n", len(cols))
		values := make([]interface{}, len(cols))
		valuePointers := make([]interface{}, len(cols))

		for i := range cols {
			valuePointers[i] = &values[i]
		}

		fmt.Println("\n", valuePointers)

		if err := rows.Scan(valuePointers...); err != nil {
			return nil, err
		}

		for i, col := range cols {

			var query_value interface{}

			val := values[i]
			b, ok := val.([]byte)

			if ok {
				query_value = string(b)
			} else {
				query_value = val
			}
			r[col] = query_value

		}
		fmt.Println(r)
		query_results = append(query_results, r)

		// fmt.Println("\n resutls---------")
		// fmt.Println(query_results)
		// fmt.Println("\n", "fuck", "you", "+++++++++++++++++")

	}

	if len(query_results) < 1 {
		err = errors.New("ot found")
		return nil, err
	}

	// defer db.Close()
	return query_results, nil
}
