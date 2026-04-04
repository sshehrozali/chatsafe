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
|------|---------|--------|
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
