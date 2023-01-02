package main

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	b := getMessage()
	var m map[string]any

	err := json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	spew.Dump(m) // Good result. Because we use any, it parses all the different fields automatically.
}

func getMessage() []byte {
	return []byte( // Important gotcha: Every numeric converts to float64.
		`{
			"id":32, 
			"name":"foo"
		}`,
	)
}

// When the keys and values are uncertain
// Passing a map gives us some flexibility instead of a static stuct
