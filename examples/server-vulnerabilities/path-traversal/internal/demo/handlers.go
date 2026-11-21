package demo

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func filesHandler(root *os.Root, safe bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		var f *os.File
		var err error
		if safe {
			f, err = root.Open(name)
		} else {
			// INTENTIONALLY VULNERABLE: cleaning a path does not confine it.
			f, err = os.Open(filepath.Join("data/public", name))
		}
		if err != nil {
			http.Error(w, "file denied or absent", 403)
			return
		}
		defer f.Close()
		info, err := f.Stat()
		if err != nil || !info.Mode().IsRegular() {
			http.Error(w, "not a regular file", 403)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		io.Copy(w, io.LimitReader(f, 65536))
	}
}
