package db

import (
	"fmt"

	"github.com/darkLord19/dfon/parser"
)

// Column represents a column in table in database
type Column struct {
	TableName  string
	ColumnName string
	DataType   string
	MinVal     int64
	MaxVal     int64
	CurrentVal int64
}

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

// GetColumsAboveThreshold get all tables which are having any auto increment field above threshold
func GetColumsAboveThreshold(dbase *parser.Database) ([]Column, error) {
	if dbase.Connection == nil {
		if e := Connect(dbase); e != nil {
			return nil, e
		}
	}
	var op []Column
	switch dbase.Type {
	case "postgres":
		getAuroIncrFiledsPsql(dbase)
	case "mysql":
		fields, err := getAuroIncrFiledsMsql(dbase)
		if err != nil {
			return nil, fmt.Errorf("Failed to retrieve auto incremating columns")
		}
		for _, field := range fields {
			if isAboveThreshold(field, dbase.Threshold) {
				op = append(op, field)
			}
		}
	}
	return op, nil
}

func isAboveThreshold(col Column, thresh float64) bool {
	if ((100 * float64(col.CurrentVal)) / float64(col.MaxVal)) >= thresh {
		return true
	}
	return false
}
