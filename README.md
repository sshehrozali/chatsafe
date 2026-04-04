# Chatsafe

**One command → one timestamped backup** of your Cursor **User** folder (`workspaceStorage`, `globalStorage`, `History`, settings).  
Quit Cursor first so files aren’t locked.

---

### Install

| Way | Command |
|-----|---------|
| **Prebuilt** | Grab the archive for your OS from [**Releases**](https://github.com/sshehrozali/chatsafe/releases), unpack, run `chatsafe` (or `chatsafe.exe` on Windows). |
| **Go** | `go install github.com/sshehrozali/chatsafe/cmd/chatsafe@latest` |

Verify: `chatsafe -version`

---

### Use

```bash
chatsafe
```

Writes `backup/cursor-backup_YYYYMMDD_HHMMSS.tar.gz` (folder is created automatically).

| Flag | Purpose |
|------|---------|
| `-out dir` | Where to put archives (default: `backup`) |
| `-cursor-user path` | Cursor **User** directory if yours isn’t in the default place |
| `-version` | Print the build tag |

**Maintainers:** tag a release to build binaries in CI:

```bash
make release-ci TAG=v1.0.1
```

---

### Develop

```bash
go test ./...
```
