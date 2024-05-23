package di

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import the pq driver
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "Postgres2024!"
	dbname   = "gocrud"
)

func providePostgresClient() (*sql.DB, func(), error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Connection succeed!")

	cleanup := func() {
		_ = db.Close()
	}

	return db, cleanup, nil
}
