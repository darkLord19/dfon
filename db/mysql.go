package db

import (
	"database/sql"

	"github.com/darkLord19/dfon/parser"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
)

// ConnectMySQL to MySQL database and assigns
func ConnectMySQL(db *parser.Database) error {
	dbC, err := sql.Open("mysql", "theUser:thePassword@/theDbName")
	if err != nil {
		return err
	}
	db.Connection = dbC
	return nil
}
