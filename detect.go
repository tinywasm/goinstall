package goinstall

import (
	"os/exec"
	"runtime"
	"strings"
)

func (g *Goinstall) IsInstalled() bool {
	v, err := g.GetVersion()
	if err != nil {
		return false
	}
	// g.version is like "1.25.2"
	// v is like "go version go1.25.2 linux/amd64"
	return strings.Contains(v, "go"+g.version)
}

func (g *Goinstall) GetVersion() (string, error) {
	path, err := g.GetPath()
	if err != nil {
		return "", err
	}
	cmd := exec.Command(path, "version")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func (g *Goinstall) GetPath() (string, error) {
	path, err := exec.LookPath("go")
	if err == nil {
		return path, nil
	}

	// Fallback to default installation paths
	var defaultPath string
	if runtime.GOOS == "windows" {
		defaultPath = `C:\Program Files\Go\bin\go.exe`
	} else {
		defaultPath = "/usr/local/go/bin/go"
	}

	if _, err := exec.LookPath(defaultPath); err == nil {
		return defaultPath, nil
	}

	return "", err
}

func IsInstalled(opts ...Option) bool {
	return New(opts...).IsInstalled()
}

func GetVersion(opts ...Option) (string, error) {
	return New(opts...).GetVersion()
}
