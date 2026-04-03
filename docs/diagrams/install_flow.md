# Install Flow

```mermaid
flowchart TD
    A[EnsureInstalled] --> B{go found\nin PATH?}
    B -->|no| INST[Install]
    B -->|yes| C{version ==\nrequired?}
    C -->|yes| AFTER
    C -->|no| RM[removeExisting]
    RM --> INST

    INST --> GOOS{runtime.GOOS?}

    GOOS -->|linux / darwin| SH[Run embedded\ninstall.sh]
    SH --> DETECT[uname -s → linux/darwin\nuname -m → amd64/arm64]
    DETECT --> DL[Download\ngo1.x.y.OS-ARCH.tar.gz\nfrom go.dev/dl/]
    DL --> RMG[rm -rf /usr/local/go]
    RMG --> TAR[tar -C /usr/local -xzf go.tar.gz]
    TAR --> LINK[ln -sf /usr/local/go/bin/go\n/usr/local/bin/go]
    LINK --> VER

    GOOS -->|windows| PS[Run embedded\ninstall.ps1]
    PS --> DL_MSI[Download\ngo1.x.y.windows-amd64.msi\nfrom go.dev/dl/]
    DL_MSI --> MSI[msiexec /i go.msi\n/quiet /norestart]
    MSI --> VER

    VER{go version\n== required?}
    VER -->|no| ERR[Cleanup + return error]
    VER -->|yes| AFTER{AfterInstall\ncallback set?}
    AFTER -->|no| RET[Return goPath]
    AFTER -->|yes| CB[Run AfterInstall\ne.g. tinygo.EnsureInstalled]
    CB -->|error| ERR
    CB -->|ok| RET
```

## removeExisting

| Platform | Action |
|----------|--------|
| Linux / macOS | `rm -rf /usr/local/go` + `rm /usr/local/bin/go` |
| Windows | `msiexec /x {ProductCode} /quiet` or `Remove-Item C:\Go -Recurse` |

## Note on shell hash cache

After install the user's shell may have the old binary path cached.
The CLI prints `hash -r` as a reminder — cannot be automated from a subprocess.
