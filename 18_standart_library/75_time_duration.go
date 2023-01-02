package main

import "time"

func main(){
	ticker := time.NewTicker(1000) // time.Durations is alias of int64, represents nanoseconds
	for {
		select {
		case <- ticker.C: // Every millisecond
			// Do smthng
		}
	}
}