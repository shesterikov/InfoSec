package demo

import (
	"net/http"
)

func New() (http.Handler, func(), error) {
	mux := http.NewServeMux()
	for _, mode := range []string{"vulnerable", "safe"} {
		mux.HandleFunc("POST /"+mode+"/xml", xmlHandler(mode == "safe"))
	}
	return mux, func() {}, nil
}
