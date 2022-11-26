package main

import (
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

func main() {
	tracing := false
	var client *http.Client

	if tracing {
		client := &http.Client{ //resolution 1
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		}

		spew.Dump(client)
	}

	if !tracing {
		client = &http.Client{ //resolution 2
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		}
		spew.Dump(client)
	}

	spew.Dump(client) //always nil in all cases !!!!
}
