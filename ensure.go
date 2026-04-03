package goinstall

import (
	"fmt"
)

func (g *Goinstall) EnsureInstalled() (string, error) {
	if g.IsInstalled() {
		return g.GetPath()
	}

	// Try to install
	if err := g.Install(); err != nil {
		return "", err
	}

	// Re-check after installation
	if !g.IsInstalled() {
		return "", fmt.Errorf("go version mismatch after installation")
	}

	goPath, err := g.GetPath()
	if err != nil {
		return "", fmt.Errorf("failed to get go path after installation: %w", err)
	}

	// Run after-install callback if provided
	if g.afterInstall != nil {
		g.log("Running after-install callback...")
		if err := g.afterInstall(goPath); err != nil {
			return goPath, fmt.Errorf("after-install callback failed: %w", err)
		}
	}

	return goPath, nil
}

func EnsureInstalled(opts ...Option) (string, error) {
	return New(opts...).EnsureInstalled()
}
