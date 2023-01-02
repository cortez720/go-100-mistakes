package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event1 struct {
	ID        int
	time.Time // time.Time implement the Marhsaler interface and change the default marhsaling behavior as embedded field.
	// Same thing with json.Unmarshal
}

type Event2 struct { // Solution 1
	ID   int
	Time time.Time // Fix with named field.
}

func (e Event1) MarshalJSON() ([]byte, error) { // Solution 2 reimplement Marshaler interface.
	return json.Marshal( // But we should keep this func up to date with current Event1 struct.
		struct {
			ID   int
			Time time.Time
		}{
			ID:   e.ID,
			Time: e.Time,
		},
	)
}

func main() {
	ev := Event1{123, time.Now()}
	bytes, _ := json.Marshal(ev)
	fmt.Print(string(bytes)) // Just "2023-01-02T23:26:07.518225216+05:00", No ID in stuct
}
