# Chatsafe

Back up Cursor to a dated archive on disk—one command, one `.tar.gz` in a folder you choose.

**Before you run:** quit Cursor so files aren’t locked.

## Install

- **Binary:** [Releases](https://github.com/sshehrozali/chatsafe/releases) → download for your OS, unpack, run `chatsafe` (or `chatsafe.exe` on Windows).
- **Go:** `go install github.com/sshehrozali/chatsafe/cmd/chatsafe@latest`

Check the build: `chatsafe -version`

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
