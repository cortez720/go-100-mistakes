package main

import (
	"database/sql"
	"errors"
)

func main() {}

func listing1() {
	if err := query(); err != nil {
		if err == sql.ErrNoRows { // Check only last error.
			// If sql.ErrNoRows is wrapped using fmt.Errorf("..: %w", err) or whatever,
			// err == sql.ErrNoRows will be always false.
			//....
		} else {
			//..
		}
	}

}

func listing2() {
	if err := query(); err != nil {
		if errors.Is(sql.ErrNoRows, err) { // Check the entire errors chain.
			// Work even it's wrapped using fmt.Errof("...%w: ", err)
			//....
		} else {
			//..
		}
	}

}

func query() error { return nil }
