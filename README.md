# goinstall
<img src="docs/img/badges.svg">

Automated Go installer — unattended, cross-platform, version-pinned (`1.25.2` by default).

Installs Go following the official instructions at https://go.dev/doc/install

---

## Quick Start — One command, no git required

### Linux / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/tinywasm/goinstall/main/scripts/install.sh | sudo bash -s 1.25.2
```

### Windows (PowerShell as Administrator)

```powershell
$env:GO_VERSION="1.25.2"; irm https://raw.githubusercontent.com/tinywasm/goinstall/main/scripts/install.ps1 | iex
```

### After install

```bash
hash -r        # Linux/macOS: clear shell cache (once per terminal)
go version     # expected: go version go1.25.2
```

> On Windows, open a **new** PowerShell window — the MSI updates PATH in the registry
> and the current window does not pick up the change.

---

## What each script does

### install.sh (Linux + macOS)

1. Detects OS (`linux` / `darwin`) and architecture (`amd64` / `arm64`)
2. Downloads `go1.25.2.OS-ARCH.tar.gz` from `https://go.dev/dl/`
3. Removes any existing `/usr/local/go`
4. Extracts the archive to `/usr/local`
5. Creates symlink: `/usr/local/bin/go` -> `/usr/local/go/bin/go`
6. Verifies with `go version`
7. Deletes the downloaded archive

### install.ps1 (Windows)

1. Downloads `go1.25.2.windows-amd64.msi` from `https://go.dev/dl/`
2. Runs the MSI installer silently (`/quiet /norestart`)
3. Verifies with `go version`
4. Deletes the downloaded MSI

---

## Alternative: clone and run locally

If you prefer to inspect the script before running:

```bash
git clone https://github.com/tinywasm/goinstall.git
cd goinstall
# Linux/macOS
sudo bash scripts/install.sh
# Windows (PowerShell as Admin)
.\scripts\install.ps1
```

---

## Programmatic use (Go API)

Use `goinstall` from another Go package to ensure Go is installed before running further setup:

```go
import "github.com/tinywasm/goinstall"

// Ensure Go 1.25.2 is installed
goPath, err := goinstall.EnsureInstalled()

// With options
goPath, err := goinstall.EnsureInstalled(
    goinstall.WithVersion("1.25.2"),
    goinstall.WithLogger(func(msg string) { fmt.Println(msg) }),
)

// Chain another installer after Go is ready
goPath, err := goinstall.EnsureInstalled(
    goinstall.WithAfterInstall(func(goPath string) error {
        _, err := tinygo.EnsureInstalled()
        return err
    }),
)
```

### Available functions

```go
goinstall.EnsureInstalled(opts ...Option) (string, error) // install if needed, return go binary path
goinstall.Install(opts ...Option) error                   // force install
goinstall.GetVersion(opts ...Option) (string, error)      // e.g. "go version go1.25.2 linux/amd64"
goinstall.IsInstalled(opts ...Option) bool                // true if required version is present
goinstall.GetPath(opts ...Option) (string, error)         // path to go binary
```

### Available options

```go
goinstall.WithVersion("1.25.2")              // pin a specific version
goinstall.WithLogger(func(msg string) {...}) // receive progress messages
goinstall.WithAfterInstall(func(goPath string) error {...}) // run after install completes
```

---

## Permissions

| Platform | Requires | Why |
|----------|----------|-----|
| Linux / macOS | `sudo` | Installs to `/usr/local` |
| Windows | Administrator PowerShell | MSI needs admin rights |

---

## Troubleshooting

**`go: command not found` after install (Linux/macOS)**
```bash
hash -r
go version
```

**`go: command not found` after install (Windows)**

Open a new PowerShell or terminal window.

**`permission denied` on Linux/macOS**

```bash
curl -fsSL ... | sudo bash    # note: sudo goes before bash, not before curl
```

**Wrong version shown**

The shell is using a cached or different Go:
```bash
which go          # Linux/macOS
where go          # Windows
```
