# Chatsafe

Back up Cursor to a dated archive on disk—one command, one `.tar.gz` in a folder you choose.

**Before you run:** quit Cursor so files aren’t locked.

## Install

Pick **your** OS/CPU, `cd` to the folder where you want the binary, then run the matching command. Each one downloads the **latest** release, unpacks **in that folder**, removes the archive, and runs `./chatsafe` once (add flags at the end if you want).

### macOS (Apple Silicon, M1/M2/M3)

```bash
curl -fsSL -o chatsafe.tgz "https://github.com/sshehrozali/chatsafe/releases/latest/download/chatsafe-darwin-arm64.tar.gz" \
  && tar -xzf chatsafe.tgz && rm chatsafe.tgz \
  && ./chatsafe
```

### macOS (Intel)

```bash
curl -fsSL -o chatsafe.tgz "https://github.com/sshehrozali/chatsafe/releases/latest/download/chatsafe-darwin-amd64.tar.gz" \
  && tar -xzf chatsafe.tgz && rm chatsafe.tgz \
  && ./chatsafe
```

### Linux (x86_64)

```bash
curl -fsSL -o chatsafe.tgz "https://github.com/sshehrozali/chatsafe/releases/latest/download/chatsafe-linux-amd64.tar.gz" \
  && tar -xzf chatsafe.tgz && rm chatsafe.tgz \
  && ./chatsafe
```

### Linux (arm64)

```bash
curl -fsSL -o chatsafe.tgz "https://github.com/sshehrozali/chatsafe/releases/latest/download/chatsafe-linux-arm64.tar.gz" \
  && tar -xzf chatsafe.tgz && rm chatsafe.tgz \
  && ./chatsafe
```

### Windows (x64, PowerShell)

```powershell
$base = "https://github.com/sshehrozali/chatsafe/releases/latest/download"
Invoke-WebRequest -Uri "$base/chatsafe-windows-amd64.zip" -OutFile ".\chatsafe.zip"
Expand-Archive -Path ".\chatsafe.zip" -DestinationPath "." -Force
Remove-Item ".\chatsafe.zip"
.\chatsafe.exe
```

Needs `curl` and `tar` on macOS/Linux. Move `chatsafe` / `chatsafe.exe` onto your `PATH` if you want it available everywhere.

**Other ways:** [all releases](https://github.com/sshehrozali/chatsafe/releases) · `go install github.com/sshehrozali/chatsafe/cmd/chatsafe@latest`

## Usage

```bash
chatsafe
```

Creates `backup/cursor-backup_YYYYMMDD_HHMMSS.tar.gz` (the `backup` folder is created if needed).

| Flag | What it does |
|------|----------------|
| `-out dir` | Where to save archives (default: `backup`) |
| `-cursor-user path` | Override if Cursor’s data isn’t in the usual place on your system |
| `-version` | Print version |

**Release maintainers:** push a new tag to build binaries in CI:

```bash
make release-ci TAG=v1.0.1
```

## Develop

```bash
go test ./...
```
