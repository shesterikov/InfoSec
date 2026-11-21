package demo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExternalEntityBoundary(t *testing.T) {
	name := filepath.Join(t.TempDir(), "secret.txt")
	if err := os.WriteFile(name, []byte("TEST_SECRET"), 0600); err != nil {
		t.Fatal(err)
	}
	data := []byte(fmt.Sprintf(`<!DOCTYPE message [<!ENTITY e SYSTEM "file://%s">]><message><value>&e;</value></message>`, name))
	got, err := parseVulnerable(data)
	if err != nil || !strings.Contains(got, "TEST_SECRET") {
		t.Fatalf("XXE: %q %v", got, err)
	}
	if _, err := parseSafe(data); err == nil {
		t.Fatal("safe parser accepted DTD")
	}
	got, err = parseSafe([]byte(`<message><value>Hello &amp; XML</value></message>`))
	if err != nil || got != "Hello & XML" {
		t.Fatalf("normal XML: %q %v", got, err)
	}
}
