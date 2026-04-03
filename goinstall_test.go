package goinstall

import (
	"os"
	"os/exec"
	"testing"
)

func TestGetVersion(t *testing.T) {
	// Ensure Go is available for the test
	_, err := exec.LookPath("go")
	if err != nil {
		t.Skip("Go is not installed, skipping TestGetVersion")
	}

	g := New()
	version, err := g.GetVersion()
	if err != nil {
		t.Fatalf("GetVersion failed: %v", err)
	}
	if version == "" {
		t.Fatal("GetVersion returned empty string")
	}
	t.Logf("Detected Go version: %s", version)
}

func TestGetPath(t *testing.T) {
	g := New()
	path, err := g.GetPath()
	if err != nil {
		t.Fatalf("GetPath failed: %v", err)
	}
	if path == "" {
		t.Fatal("GetPath returned empty string")
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("Go path does not exist: %s", path)
	}
	t.Logf("Detected Go path: %s", path)
}
