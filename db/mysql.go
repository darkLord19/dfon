package db

import (
	"database/sql"

	"github.com/darkLord19/dfon/parser"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
)

func connectMySQL(dbase *parser.Database) error {
	dbC, err := sql.Open("mysql", "theUser:thePassword@/theDbName")
	if err != nil {
		return err
	}
	dbase.Connection = dbC
	return nil
}
