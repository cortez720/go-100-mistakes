package main

import (
	"database/sql"

	"github.com/google/uuid"
)

func main() {
	var db *sql.DB
	ids := make([]uuid.UUID, 0)
	// Statmenet doesn't have to be recompiled (parsing + optimization + translation)
	// Redces the risks of sql injections
	stmt, err := db.Prepare("SELECT * FROM LOCATIONS WHERE uuid = ?")
	if err != nil {
		panic(err)
	}

	var rows *sql.Rows
	for id := range ids {
		rows, err = stmt.Query(id) // Use 1 prepared statement many times. // Could use it councurrently.
		rows.Scan()
	}

	stmt.Close() // Close statement after use

}

// We need to remember to use prepared statements when it makes sense.
