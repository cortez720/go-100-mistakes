// err := rows.Next can break either when there are no more rows or when an  error happens
// Following rows iteration we shuld call rows.Err

package main

import "database/sql"

func main() {
	var rows sql.Rows
	var department string
	var age int

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	for rows.Next() { // Can stop because of the error.
		err := rows.Scan(&department, age)
		if err != nil { // Not enough
			panic(err)
		}
	}

	if err := rows.Err(); err != nil {
		panic(err) // We should keep this in mind
	}
}
