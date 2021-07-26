package main

func main() {
	db_type := "postgres"
	db_user := "postgres"
	db_password := "123456"
	db_host := "localhost"
	db_database := "dbSapHasap"

	a := App{}
	a.Initialize(
		db_type,
		db_user,
		db_password,
		db_host,
		db_database,
	)

	a.Run(":8080")

}
