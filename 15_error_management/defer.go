package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const rawQuery = "..."

func getBalance1(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(rawQuery, clientID)
	if err != nil {
		return 0, err
	}
	defer rows.Close() // It's ignored unintentionally

	// Use rows
	return 0, nil
}

func getBalance2(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(rawQuery, clientID)
	if err != nil {
		return 0, err
	}
	defer func() { _ = rows.Close() }() // If we dont need the error, just ignore it.

	// Use rows
	return 0, nil
}

func getBalance3(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(rawQuery, clientID)
	if err != nil {
		return 0, err
	}
	defer func() {
		err = rows.Close()
	}() // Not what we needed. It overwrites the db.Query err. It always returns rows.Close() error.

	// Use rows
	return 0, nil
}

func getBalance4(db *sql.DB, clientID string) (balance float32, err error) {
	rows, err := db.Query(rawQuery, clientID)
	if err != nil {
		return 0, err
	}
	defer func() {
		closeErr := rows.Close()
		if err != nil { // If rows.Scan error != nil, we just log rows.Close error (if it not nil).
			if closeErr != nil {
				log.Printf("failed to close rows: %v", err)
			}
			return // Here we return rows.Scan error, whatever nil rows.Close or not.
		}

		err = closeErr // return rows.Close error if rows.Scan is nil. Here error can be either nil or error.
	}()

	err = rows.Scan()

	// Use rows
	return 0, nil
}

var ErrObjRowsCloseAndScan = errors.New("rows scanning and rows closing are failed")

func getBalance5(db *sql.DB, clientID string) (balance float32, err error) {
	rows, err := db.Query(rawQuery, clientID)
	if err != nil {
		return 0, err
	}
	defer func() {
		closeErr := rows.Close()
		if err != nil {
			if closeErr != nil {
				err = ErrObjRowsCloseAndScan // Assign third special error object ad hoc.
			}
			return
		}

		err = closeErr // return rows.Close error if rows.Scan is nil. Here error can be either nil or error.
	}()

	err = rows.Scan()

	// Use rows
	return 0, nil
}

func main() {
	err := errors.New("Scan error")
	closeErr := errors.New("Close error")
	err = fmt.Errorf("...%w...:%v...:%v", ErrObjRowsCloseAndScan, err, closeErr)
	// Pass all errors through, errors.Is only with ErrObj..

	fmt.Println(err, errors.Is(err, ErrObjRowsCloseAndScan))

	err = fmt.Errorf("%w")
}
