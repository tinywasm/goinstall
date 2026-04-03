# goinstall

Automated Go installer — unattended, cross-platform, version-pinned (`1.25.2` by default).

Installs Go following the official instructions at https://go.dev/doc/install

---

## Quick Start — Which script do I run?

| My system | Script to use |
|-----------|---------------|
| Linux (Ubuntu, Debian, Arch…) | `scripts/install.sh` |
| macOS | `scripts/install.sh` |
| Windows | `scripts/install.ps1` |

---

## Linux / macOS — Step by step

### 1. Clone or download the repo

```bash
git clone https://github.com/tinywasm/goinstall.git
cd goinstall
```

### 2. Give the script execute permission

```bash
chmod +x scripts/install.sh
```

### 3. Run with sudo (required — installs to /usr/local)

```bash
sudo bash scripts/install.sh
```

To install a specific version:

```bash
sudo bash scripts/install.sh 1.25.2
```

### 4. Verify

```bash
hash -r          # clear shell cache (run this once in the same terminal)
go version       # expected: go version go1.25.2 linux/amd64
```

> **Why `hash -r`?** If you had an older Go installed, your shell cached the old path.
> `hash -r` tells the shell to look up commands again. Only needed once per terminal session.

### What the script does

1. Detects your OS (`linux` or `darwin`) and architecture (`amd64` or `arm64`)
2. Downloads `go1.25.2.OS-ARCH.tar.gz` from `https://go.dev/dl/`
3. Removes any existing `/usr/local/go`
4. Extracts the archive to `/usr/local`
5. Creates a symlink: `/usr/local/bin/go → /usr/local/go/bin/go`
6. Verifies with `go version`
7. Deletes the downloaded archive

---

## Windows — Step by step

### 1. Open PowerShell as Administrator

Press `Win + X` → select **Windows PowerShell (Admin)** or **Terminal (Admin)**.

### 2. Allow script execution (one-time setup)

```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

### 3. Clone or download the repo

```powershell
git clone https://github.com/tinywasm/goinstall.git
cd goinstall
```

### 4. Run the script

```powershell
.\scripts\install.ps1
```

To install a specific version:

```powershell
.\scripts\install.ps1 1.25.2
```

### 5. Verify

Open a **new** PowerShell window, then:

```powershell
go version    # expected: go version go1.25.2 windows/amd64
```

> **Why a new window?** The Windows MSI installer updates the PATH in the registry.
> The current window does not pick up the change — a new window does.

### What the script does

1. Downloads `go1.25.2.windows-amd64.msi` from `https://go.dev/dl/`
2. Runs the MSI installer silently (`/quiet /norestart`)
3. Verifies with `go version`
4. Deletes the downloaded MSI

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

| Platform | Requires |
|----------|----------|
| Linux / macOS | `sudo` — installs to `/usr/local` |
| Windows | Administrator PowerShell — MSI needs admin rights |

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
Make sure you ran the script with `sudo`:
```bash
sudo bash scripts/install.sh
```

**Wrong version shown**
The shell is using a cached or different Go. Check which one is active:
```bash
which go          # Linux/macOS
where go          # Windows
```
