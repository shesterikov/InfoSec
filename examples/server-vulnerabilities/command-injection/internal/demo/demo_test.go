package demo

import (
	"context"
	"strings"
	"testing"
)

func TestShellBoundary(t *testing.T) {
	payload := "Alice'; id; #"
	unsafe, err := command(context.Background(), payload, false)
	if err != nil || !strings.Contains(string(unsafe), "uid=") {
		t.Fatalf("shell demo: %q %v", unsafe, err)
	}
	safe, err := command(context.Background(), payload, true)
	if err != nil || string(safe) != payload+"\n" {
		t.Fatalf("argument boundary: %q %v", safe, err)
	}
}
