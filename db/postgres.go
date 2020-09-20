package db

import (
	"database/sql"

	"github.com/darkLord19/dfon/parser"
	_ "github.com/lib/pq" // import postgres driver
)

func connectPostgres(db *parser.Database) error {
	dbC, err := sql.Open("postgres", "user=theUser dbname=theDbName sslmode=verify-full")
	if err != nil {
		return err
	}
	db.Connection = dbC
	return nil
}
