package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event3 struct {
	Time time.Time
}

func main() {
	t := time.Now()
	ev1 := Event3{t}

	bytes, _ := json.Marshal(ev1)

	var ev2 Event3 // No monolitic part of time.
	json.Unmarshal(bytes, &ev2)

	fmt.Println(ev1 == ev2)               // False. One with monolitick part, other not.
	fmt.Println(ev1.Time.Equal(ev2.Time)) // Solution 1 // True // Equal doesn't consider monolitic time.

	ev11 := Event3{t.Truncate(0)} // Solution 2 // With 0 returns without monolitic clock

	fmt.Println(ev11 == ev2) // True

	withLocation()
}

func withLocation(){ // time.Now depends on location.
	location, err := time.LoadLocation("America/New_York")
	if err != nil{
		panic(err)
	}
	fmt.Println(time.Now().In(location)) // In specific location
	fmt.Println(time.Now().UTC())  // UTC
}