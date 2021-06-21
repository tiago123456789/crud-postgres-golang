package database

import (
	"database/sql"
	"fmt"
)

func CreateConnection() *sql.DB {
	url := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	 	"127.0.0.1", 5433, "postgres", "postgres", "crud"
	)

	fmt.Println(url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("No possible connect in database")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
