package main

import "database/sql"

func main() {
	var rows sql.Rows

	defer func() {
		err := rows.Close() // We must close the connection
		// Forgetting to do this means a connecion leak, if it doesn't return an error
		if err != nil {
			panic(err)
		}
	}()

}
