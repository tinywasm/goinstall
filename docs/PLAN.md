# PLAN: goinstall — Automated Go Installer

## Goal

Unattended, cross-platform Go installation pinned to a specific version (default: `1.25.2`).
Scripts are the installation unit. The Go package embeds them and provides a programmatic API
so other packages can chain installations after Go is ready (e.g. `tinygo.EnsureInstalled()`).

## Design Decisions

- **Scripts are the source of truth**: bash (Linux+macOS) and PowerShell (Windows) do the actual work.
- **Single bash script**: detects OS (`uname -s`) and architecture (`uname -m`) — no separate Linux/macOS files.
- **Go package embeds scripts** via `embed.FS` and runs them — zero external dependencies.
- **Version-pinned**: `DefaultVersion = "1.25.2"` — explicit, no "latest" ambiguity.
- **Idempotent**: correct version already installed → return immediately.
- **Version mismatch → auto-update**: remove old install, run script again.
- **No shell modification**: symlink `/usr/local/bin/go` makes Go available without touching `.bashrc`/`.zshrc`.
- **Post-install callback**: optional `WithAfterInstall` hook to chain further installations.

## Consumer Pattern

```go
// Simple use
goPath, err := goinstall.EnsureInstalled()

// With chained installer (e.g. tinygo after go)
goPath, err := goinstall.EnsureInstalled(
    goinstall.WithVersion("1.25.2"),
    goinstall.WithLogger(func(msg string) { fmt.Println(msg) }),
    goinstall.WithAfterInstall(func(goPath string) error {
        _, err := tinygo.EnsureInstalled()
        return err
    }),
)
```

## Installation Method per Platform

| Platform | Script | Method | Admin |
|----------|--------|--------|-------|
| Linux amd64/arm64 | `scripts/install.sh` | tarball → `/usr/local/go` + symlink | `sudo` |
| macOS amd64/arm64 | `scripts/install.sh` | tarball → `/usr/local/go` + symlink | `sudo` |
| Windows amd64 | `scripts/install.ps1` | MSI silent install (`/quiet`) | Admin shell |

## Script: install.sh (Linux + macOS)

```bash
OS=$(uname -s | tr '[:upper:]' '[:lower:]')      # linux / darwin
ARCH=$(uname -m)                                  # x86_64 → amd64, aarch64 → arm64
# Download go1.25.2.$OS-$ARCH.tar.gz from https://go.dev/dl/
# rm -rf /usr/local/go
# tar -C /usr/local -xzf go.tar.gz
# ln -sf /usr/local/go/bin/go /usr/local/bin/go
# go version  ← verify
```

## Script: install.ps1 (Windows)

```powershell
# Download go1.25.2.windows-amd64.msi from https://go.dev/dl/
# msiexec /i go.msi /quiet /norestart
# go version  ← verify
```

## Public API

```go
func EnsureInstalled(opts ...Option) (string, error)  // check version → install if needed → run AfterInstall
func Install(opts ...Option) error                     // force install
func GetVersion(opts ...Option) (string, error)        // returns "go1.25.2"
func IsInstalled(opts ...Option) bool

func WithVersion(v string) Option
func WithLogger(f func(string)) Option
func WithAfterInstall(f func(goPath string) error) Option
```

## Files

| File | Role |
|------|------|
| `goinstall.go` | Config, options, constants, `//go:embed` |
| `detect.go` | `GetVersion`, `IsInstalled`, `GetPath` |
| `ensure.go` | `EnsureInstalled` — version check + AfterInstall callback |
| `install.go` | `Install` — platform dispatch, removeExisting, run script |
| `scripts/install.sh` | Bash script for Linux + macOS |
| `scripts/install.ps1` | PowerShell script for Windows |
| `cmd/goinstall/main.go` | CLI: `-version`, `-v`, `-h` flags |

## Stages

| Stage | Description | Completed |
|-------|-------------|-----------|
| 1 | `scripts/install.sh` — bash script (detect OS/arch, download, extract, symlink, verify) | [ ] |
| 2 | `scripts/install.ps1` — PowerShell script (download MSI, silent install, verify) | [ ] |
| 3 | `goinstall.go` — config, options, constants, embed | [ ] |
| 4 | `detect.go` — GetVersion, IsInstalled, GetPath | [ ] |
| 5 | `install.go` — platform dispatch, removeExisting, run embedded script | [ ] |
| 6 | `ensure.go` — EnsureInstalled + AfterInstall callback | [ ] |
| 7 | `cmd/goinstall/main.go` — CLI entry point | [ ] |

## Flow Diagram

See [docs/diagrams/install_flow.md](diagrams/install_flow.md)
