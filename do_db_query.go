package main

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"reflect"
)

type query_response map[string]interface{}

func do_db_select_query(db *sql.DB, query_string string, base64_columns []string) ([]query_response, error) {

	rows, err := db.Query(query_string)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	query_results := []query_response{}

	for rows.Next() {

		r := make(query_response)
		cols, _ := rows.Columns()

		values := make([]interface{}, len(cols))
		valuePointers := make([]interface{}, len(cols))

		for i := range cols {
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			return nil, err
		}

		for i, col := range cols {

			var query_value interface{}

			val := values[i]
			// fmt.Printf("%T\n", val)
			// fmt.Println(reflect.TypeOf(val))

			if contains_in_slice(base64_columns, col) {
				slice_type, ok := val.([]byte)
				if ok {
					val = base64.StdEncoding.EncodeToString(slice_type)
				}
				if !ok {
					if reflect.TypeOf(val).Kind() == reflect.String {
						slice_type := []byte(val.(string))
						val = base64.StdEncoding.EncodeToString(slice_type)
					}
				}
			}

			b, ok := val.([]byte)

			if ok {
				query_value = string(b)
			} else {
				query_value = val
			}
			r[col] = query_value

		}
		query_results = append(query_results, r)
	}

	if len(query_results) < 1 {
		err = errors.New("empty response")
		return nil, err
	}

	// defer db.Close()
	return query_results, nil
}

func do_db_query_exec(db *sql.DB, query_string string) error {

	_, err := db.Query(query_string)
	if err != nil {
		return err
	}

	return nil
}
