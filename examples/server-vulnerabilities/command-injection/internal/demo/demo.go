package demo

import (
	"net/http"
)

func New() (http.Handler, func(), error) {
	mux := http.NewServeMux()
	for _, mode := range []string{"vulnerable", "safe"} {
		mux.HandleFunc("GET /"+mode+"/greet", greetHandler(mode == "safe"))
	}
	return mux, func() {}, nil
}
