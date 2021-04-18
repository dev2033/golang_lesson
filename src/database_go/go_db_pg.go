package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=admin password=1234 dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into Products "+
		"(model, company, price) "+
		"values ('Samsung a20', $1, $2)",
		"Samsung", 72000)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
