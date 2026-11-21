package demo

import (
	"io"
	"net/http"
	"net/url"
)

func importHandler(client *http.Client, safe bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		target := r.URL.Query().Get("url")
		if safe {
			// Business function accepts a document ID, never a user-supplied URL.
			if r.URL.Query().Get("document") != "catalog" {
				http.Error(w, "unknown document", 400)
				return
			}
			target = "http://127.0.0.1:9090/public"
		}
		u, err := url.Parse(target)
		if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
			http.Error(w, "invalid URL", 400)
			return
		}
		req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, u.String(), nil)
		if err != nil {
			http.Error(w, "invalid request", 400)
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "fetch failed", 502)
			return
		}
		defer resp.Body.Close()
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		io.Copy(w, io.LimitReader(resp.Body, 65536))
	}
}
