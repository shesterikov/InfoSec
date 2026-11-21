package demo

import (
	"io"
	"net/http"
)

func xmlHandler(safe bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(http.MaxBytesReader(w, r.Body, 65536))
		if err != nil {
			http.Error(w, "body too large", 413)
			return
		}
		var result string
		if safe {
			result, err = parseSafe(data)
		} else {
			result, err = parseVulnerable(data)
		}
		if err != nil {
			http.Error(w, "XML rejected", 400)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		io.WriteString(w, result)
	}
}
