package postgres

import (
	"database/sql"
	"fmt"
	"os"

	//for connecting to db
	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

var (
	DbClient *sql.DB
)

func init() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	DbClient, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	// defer dbClient.Close()
	fmt.Println((DbClient))
	fmt.Println("Successfully connected!")
}
