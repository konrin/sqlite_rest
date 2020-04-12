package main

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

type (
	DB struct {
		sqlite *sql.DB
	}
	Rows []*Row
	Row map[string]*ColumnValue

	ColumnValue struct {
		Data interface{}
		Type *sql.ColumnType
	}
)

func NewDB(connection string) (*DB, error) {
	sqlite, err := sql.Open("sqlite3", connection)
	if err != nil {
		return nil, err
	}

	db := &DB{
		sqlite: sqlite,
	}

	if err = sqlite.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) Echo(echo string) (string, error) {
	row := db.sqlite.QueryRow("SELECT ?", echo)
	if row == nil {
		return "", errors.New("echo failed")
	}

	var result string

	if err := row.Scan(&result); err != nil {
		return "", err
	}

	return result, nil
}

func (db *DB) Exec(sql string, args ...interface{}) (sql.Result, error) {
	result, err := db.sqlite.Exec(sql, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (db *DB) Query(sql string, args ...interface{}) (Rows, error) {
	rows, err := db.sqlite.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, nil
	}

	defer rows.Close()

	var rowsData = make(Rows, 0, 0)

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := Row{}

		for i, col := range columns {
			row[col] = &ColumnValue{
				Type: columnTypes[i],
			}
		}

		rowsData = append(rowsData, &row)

		dataCols := make([]interface{}, len(columns))

		for i, col := range columns {
			dataCols[i] = &row[col].Data
		}

		err = rows.Scan(dataCols...)
		if err != nil {
			return nil, err
		}
	}

	if rows.Err() != nil {
		return nil, err
	}

	return rowsData, nil
}

func (db *DB) Close() error {
	return db.sqlite.Close()
}
