package postgres

import (
	"database/sql"
	"fmt"
)

var(
	DbClient *sql.DB
)

func init() {
	var err error
	connectionString := "user=postgres dbname=postgres password= host=localhost sslmode=disable"
	DbClient, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer DbClient.Close()
	fmt.Printf("\nSuccessfully connected to database!\n")
}