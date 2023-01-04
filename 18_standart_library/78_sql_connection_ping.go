package main

import (
	"context"
	"database/sql"
	"time"
)

const dsn = ""

func main() {
	db, err := sql.Open("mysql", dsn) // Doesn't check the connection, just validate the arguments without creation connection.
	if err != nil { 
		panic(err)
	}

	if err := db.Ping(); err != nil { // Check the connection here.
		// Ensure that the data source name is valid and the database is reachable.
		panic(err)
	}

	context, _ := context.WithTimeout(context.Background(), time.Minute)

	if err := db.PingContext(context); err != nil { // Conveying when the ping should be canceled or time out
		panic(err)
	}
}
