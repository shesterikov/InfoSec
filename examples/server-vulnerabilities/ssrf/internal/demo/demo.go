// Package demo runs public and private listeners in ONE process.
package demo

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func New() (http.Handler, func(), error) {
	listener, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		return nil, nil, err
	}
	private := http.NewServeMux()
	private.HandleFunc("GET /admin", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "DEMO_INTERNAL_TOKEN=classroom-only") })
	private.HandleFunc("GET /public", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Public catalog document") })
	server := &http.Server{Handler: private, ReadHeaderTimeout: 3 * time.Second}
	go server.Serve(listener)
	client := &http.Client{Timeout: 3 * time.Second, Transport: &http.Transport{Proxy: nil}, CheckRedirect: func(r *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
	mux := http.NewServeMux()
	for _, mode := range []string{"vulnerable", "safe"} {
		mux.HandleFunc("GET /"+mode+"/import", importHandler(client, mode == "safe"))
	}
	return mux, func() { server.Close(); client.CloseIdleConnections() }, nil
}
