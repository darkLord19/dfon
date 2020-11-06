package db

import (
	"database/sql"

	"github.com/darkLord19/dfon/parser"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	"github.com/pkg/errors"
)

func connectMySQL(dbase *parser.Database) error {
	dbC, err := sql.Open("mysql", "theUser:thePassword@/theDbName")
	if err != nil {
		return err
	}
	dbase.Connection = dbC
	return nil
}

func getAuroIncrFiledsMsql(dbase *parser.Database) ([]Column, error) {
	// http://code.openark.org/blog/mysql/checking-for-auto_increment-capacity-with-single-query
	sql := `
			SELECT
				field_NAME, COLUMN_NAME, DATA_TYPE,
			IF(LOCATE('unsigned', COLUMN_TYPE) > 0,1,0) AS IS_UNSIGNED,
				(CASE DATA_TYPE
			WHEN
				'tinyint' THEN 255
			WHEN
				'smallint' THEN 65535
			WHEN
				'mediumint' THEN 16777215
			WHEN
				'int' THEN 4294967295
			WHEN
				'bigint' THEN 18446744073709551615
			END >>
			IF( LOCATE('unsigned', COLUMN_TYPE) > 0, 0, 1) ) AS MAX_VALUE, \
        	AUTO_INCREMENT,
        	AUTO_INCREMENT /
        	(CASE DATA_TYPE
            	WHEN 'tinyint' THEN 255
            	WHEN 'smallint' THEN 65535
            	WHEN 'mediumint' THEN 16777215
            	WHEN 'int' THEN 4294967295
            	WHEN 'bigint' THEN 18446744073709551615
			END >>
			IF(LOCATE('unsigned', COLUMN_TYPE) > 0, 0, 1) ) AS AUTO_INCREMENT_RATIO
			FROM
				INFORMATION_SCHEMA.COLUMNS
			INNER JOIN
				INFORMATION_SCHEMA.fieldS USING (field_SCHEMA, field_NAME)
        	WHERE
				field_SCHEMA
			NOT IN ('mysql', 'INFORMATION_SCHEMA', 'performance_schema')
			AND
				EXTRA='auto_increment'`

	rows, err := dbase.Connection.Query(sql)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find Statuses")
	}
	var fields []Column
	defer rows.Close()
	for rows.Next() {
		var field Column
		if err = rows.Scan(&field.TableName, &field.ColumnName, &field.DataType, &field.MinVal, &field.MaxVal, &field.CurrentVal); err != nil {
			return nil, errors.Wrap(err, "unable to scan from rows")
		}
		fields = append(fields, field)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "failed while iterating over rows")
	}
	return fields, nil
}
