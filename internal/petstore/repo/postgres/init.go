package postgres

import (
	"database/sql"
)

var (
	dbClient *sql.DB
)

// func init() {
// 	var err error
// 	connectionString := "user=postgres dbname=postgres password= host=localhost sslmode=disable"
// 	dbClient, err = sql.Open("postgres", connectionString)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer dbClient.Close()
// 	fmt.Printf("\nSuccessfully connected to database!\n")
// }
