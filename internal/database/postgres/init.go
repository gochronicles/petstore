package postgres

import (
	"database/sql"
	"fmt"
)

var (
	DbClient *sql.DB
)

func init() {
	var err error
	connectionString := "user=database dbname=database password= host=localhost sslmode=disable"
	DbClient, err = sql.Open("database", connectionString)
	if err != nil {
		panic(err)
	}
	defer DbClient.Close()
	fmt.Printf("\nSuccessfully connected to database!\n")
}
