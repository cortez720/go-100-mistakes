package main

import (
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

func main() {
	tracing := true
	var client *http.Client

	if tracing {
		c := &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		}
		client = c
		spew.Dump(client)
	}

	if !tracing {
		c := &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		}
		client = c
		spew.Dump(client)
	}

	spew.Dump(client) //always nil in all cases !!!!
}
