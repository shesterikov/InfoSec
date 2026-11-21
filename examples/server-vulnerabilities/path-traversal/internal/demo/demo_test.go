package demo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRootRejectsTraversalAndSymlink(t *testing.T) {
	parent := t.TempDir()
	public := filepath.Join(parent, "public")
	if err := os.Mkdir(public, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(parent, "secret"), []byte("secret"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(public, "hello"), []byte("hello"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("../secret", filepath.Join(public, "link")); err != nil {
		t.Fatal(err)
	}
	root, err := os.OpenRoot(public)
	if err != nil {
		t.Fatal(err)
	}
	defer root.Close()
	f, err := root.Open("hello")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	for _, name := range []string{"../secret", "link"} {
		f, err := root.Open(name)
		if err == nil {
			f.Close()
			t.Fatalf("escaped root using %s", name)
		}
	}
}
