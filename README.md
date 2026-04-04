# Chatsafe

Small **macOS** CLI that saves your Cursor editor data into **one `.tar.gz` file** per run. Handy before renames, OS upgrades, or whenever you want a snapshot you can stash or copy elsewhere.

## What gets backed up

Everything Chatsafe reads lives under **Cursor → User** (same idea as Cursor’s own data folder):

- `workspaceStorage` (workspaces and related state)
- `globalStorage`
- `History`
- `settings.json` and `keybindings.json` (if present)

Missing pieces are skipped; the archive only needs at least one of these to exist.

## Before you run

1. **Quit Cursor completely** (not just the window), so files are not locked.
2. **macOS only** — paths are built for the standard Cursor layout on Mac.
3. **Go 1.22+** if you use `go run` / `go build`.

## Install prebuilt binary (GitHub Releases)

Releases are published as **GitHub Releases** (versioned `.tar.gz` / `.zip` + `checksums.txt`). GitHub does not offer a separate “Packages” registry for plain CLI binaries; Releases is the usual way to ship them.

1. Open **[Releases](https://github.com/sshehrozali/chatsafe/releases)** and download the archive for your OS and CPU (e.g. macOS Apple Silicon: `chatsafe_1.0.0_darwin_arm64.tar.gz`).
2. Extract the `chatsafe` binary and run it (e.g. move it to a directory on your `PATH`).

```bash
# Example after extracting next to the binary:
./chatsafe -version
./chatsafe
```

With [GitHub CLI](https://cli.github.com/):

```bash
gh release download v1.0.0 --repo sshehrozali/chatsafe --pattern 'chatsafe_*_darwin_arm64.tar.gz'
tar -xzf chatsafe_*_darwin_arm64.tar.gz
./chatsafe -version
```

Maintainers: create a release by pushing an annotated tag; CI builds and uploads assets:

```bash
git tag -a v1.0.0 -m "chatsafe 1.0.0"
git push origin v1.0.0
```

## How to use

Clone the repo and run from its root:

```bash
git clone https://github.com/sshehrozali/chatsafe.git
cd chatsafe
go run ./cmd/chatsafe
```

You’ll see a short progress indicator, then a line like:

`backup written: backup/cursor-backup_YYYYMMDD_HHMMSS.tar.gz (… bytes)`

- **`backup/`** is created the **first** time; later runs add **new** archives next to the old ones (nothing is overwritten automatically).

### Build a binary

```bash
go build -o chatsafe ./cmd/chatsafe
./chatsafe
```

### Options

| Flag | Default | Meaning |
|------|---------|---------|
| `-version` | — | Print release tag (or `dev` when built locally) and exit. |
| `-out` | `backup` | Folder for `.tar.gz` files (relative to where you run the command). |
| `-cursor-user` | `~/Library/Application Support/Cursor/User` | Override if your Cursor data lives somewhere else. |

Example:

```bash
go run ./cmd/chatsafe -out ~/Desktop/cursor-snapshots
```

## Restoring (quick note)

Unpack the archive wherever you like and copy folders/files back into your Cursor **User** directory only if you know what you’re doing—ideally with Cursor **quit**, and preferably after backing up the current User folder first.

## Develop

```bash
go test ./...
```
