package main

import "database/sql"

func main() {
	rows := sql.Rows{}
	var (
		department string
		age        int
	)

	rows.Scan(&department, &age) // Error: Converting NULL to string is unsupported

	// Solution 1. Pointer
	var department1 *string
	rows.Scan(&department1, &age) // department1 here is nil.

	// Solution 2. sql.NullXXX
	var department2 sql.NullString
	rows.Scan(&department2, &age)

	// Both are work.
}
