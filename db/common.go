package db

import (
	"github.com/darkLord19/dfon/parser"
)

// Connect connects to a database based on it's type
// returns error if connection fails
func Connect(dbase *parser.Database) error {
	if dbase.Type == "postgres" {
		return connectPostgres(dbase)
	} else if dbase.Type == "mysql" {
		return connectMySQL(dbase)
	}
	return nil
}

// Disconnect disconnects from given database
func Disconnect(dbase *parser.Database) {
	if dbase.Connection != nil {
		dbase.Connection.Close()
	}
}

// GetTables present in database
func GetTables(dbase *parser.Database) ([]string, error) {
	query := "show tables"
	if dbase.Connection == nil {
		if e := Connect(dbase); e != nil {
			return nil, e
		}
	}
	res, err := dbase.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	var tables []string
	for res.Next() {
		var table string
		res.Scan(&table)
		tables = append(tables, table)
	}
	return tables, nil
}
