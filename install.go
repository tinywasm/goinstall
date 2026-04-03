package goinstall

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func (g *Goinstall) Install() error {
	g.log(fmt.Sprintf("Installing Go %s...", g.version))

	// Write embedded script to temporary file
	scriptName := "install.sh"
	if runtime.GOOS == "windows" {
		scriptName = "install.ps1"
	}

	scriptPath := filepath.Join("scripts", scriptName)
	scriptContent, err := scriptsFS.Open(scriptPath)
	if err != nil {
		return fmt.Errorf("failed to open embedded script %s: %w", scriptPath, err)
	}
	defer scriptContent.Close()

	tempDir, err := os.MkdirTemp("", "goinstall")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir)

	tempScriptPath := filepath.Join(tempDir, scriptName)
	tempScriptFile, err := os.Create(tempScriptPath)
	if err != nil {
		return fmt.Errorf("failed to create temp script file: %w", err)
	}

	if _, err := io.Copy(tempScriptFile, scriptContent); err != nil {
		tempScriptFile.Close()
		return fmt.Errorf("failed to write temp script file: %w", err)
	}
	tempScriptFile.Close()

	if runtime.GOOS != "windows" {
		if err := os.Chmod(tempScriptPath, 0755); err != nil {
			return fmt.Errorf("failed to set script permissions: %w", err)
		}
	}

	// Run the script
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", tempScriptPath, g.version)
	} else {
		cmd = exec.Command("bash", tempScriptPath, g.version)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("installation script failed: %w", err)
	}

	g.log("Go installed successfully.")
	return nil
}

func Install(opts ...Option) error {
	return New(opts...).Install()
}

