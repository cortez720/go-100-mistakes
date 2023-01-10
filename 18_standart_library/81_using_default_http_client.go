package main

import (
	"net"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second, // Time limit for a request. Includes all steps, from dial to read response body.
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second, // Maximum amount of time a dial will wait for a connection to complete
			}).DialContext,
			TLSHandshakeTimeout:   time.Second, // Maximum time of TLS handshake
			ResponseHeaderTimeout: time.Second, // Amount of time to wait for a server's response headers.
		},
	}
}

// By default, HTTP client doews connection pooling. The default client reuses connections (it can be disabled by setting http.Transport.DisableKeepAlives to true)
// extra tomeout to specify how long an idle  connecition is kep in  the pool: http.Transport.IdleConneTimeout default: 90s.
// to configure the number of connections in the pool,, we must override http.Transport.MaxIdleConns. Default: 100. But http.Transport.MaxIdleConnsOerHost limit per host. Default :2. We will have to repoen at least 98 connections.
// For production-grade systems we probably want  to override the default timeouts.
// And tweaking the parameters related to connection pooling cn also have a significant impact on the latency.
