package demo

import (
	"context"
	"net/http"
	"os/exec"
	"time"
)

func command(ctx context.Context, value string, safe bool) ([]byte, error) {
	if safe {
		return exec.CommandContext(ctx, "/usr/bin/printf", "%s\n", value).CombinedOutput()
	}
	// INTENTIONALLY VULNERABLE: shell parses metacharacters in value.
	return exec.CommandContext(ctx, "/bin/sh", "-c", "printf '%s\\n' '"+value+"'").CombinedOutput()
}

func greetHandler(safe bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		value := r.URL.Query().Get("name")
		if len(value) > 1024 {
			http.Error(w, "name too long", 400)
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		out, err := command(ctx, value, safe)
		if err != nil {
			http.Error(w, "command failed", 400)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write(out)
	}
}
