package demo

import (
	"net/http"
	"os"
)

func New() (http.Handler, func(), error) {
	root, err := os.OpenRoot("data/public")
	if err != nil {
		return nil, nil, err
	}
	mux := http.NewServeMux()
	for _, mode := range []string{"vulnerable", "safe"} {
		mux.HandleFunc("GET /"+mode+"/files", filesHandler(root, mode == "safe"))
	}
	return mux, func() { root.Close() }, nil
}
