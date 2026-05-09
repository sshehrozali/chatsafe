# Chatsafe

Back up Cursor to a dated archive on disk—one command, one `.tar.gz` in a folder you choose.

**Portable:** install once with the command below; **`chatsafe`** lands on your **`PATH`** so you can run it from **any directory**.

**Before you run:** quit Cursor so files aren’t locked.

<h3 align="center">🚀 <strong>Claude support upcoming</strong> — <em>coming soon</em> 🤖 🔜</h3>

## Install (one command)

**macOS / Linux** — installs to `~/.local/bin` and adds it to your shell config (`~/.zshrc`, `~/.bashrc`, or `~/.profile`):

```bash
curl -fsSL https://raw.githubusercontent.com/sshehrozali/chatsafe/main/install.sh | sh
```

**Windows (PowerShell)** — installs to `%USERPROFILE%\bin` and updates your user `PATH`:

```powershell
irm https://raw.githubusercontent.com/sshehrozali/chatsafe/main/install.ps1 | iex
```

Reload your shell config (same window) **or** open a new terminal, then check:

```bash
source ~/.zshrc   # zsh — use ~/.bashrc if you use bash
chatsafe -version
```

Until you do that, `chatsafe` is not on `PATH` in this session. You can always run `~/.local/bin/chatsafe -version` right away.

Windows: open a **new** PowerShell, then `chatsafe.exe -version`.

Optional custom folder (macOS/Linux):  
`curl -fsSL https://raw.githubusercontent.com/sshehrozali/chatsafe/main/install.sh | env BIN_DIR="$HOME/my/bin" sh`

**Other ways:** [all releases](https://github.com/sshehrozali/chatsafe/releases) · `go install github.com/sshehrozali/chatsafe/cmd/chatsafe@latest`

## Usage

Backups run only when you pass the **`backup`** subcommand:

```bash
chatsafe backup
```

Creates `backup/cursor-backup_YYYYMMDD_HHMMSS.tar.gz` in your **current working directory** (the `backup` folder is created if needed).

### Examples

```bash
# Default output directory under your current directory
cd ~
chatsafe backup
```

```bash
# Save archives somewhere fixed (folder is created if needed)
chatsafe backup -out ~/Documents/cursor-backups
```

```bash
# Cursor data lives somewhere non-standard—point at the User folder explicitly
chatsafe backup -cursor-user "/path/to/Cursor/User"
```

| Command / flag | What it does |
|----------------|--------------|
| `chatsafe backup` | Create a timestamped archive |
| `-out dir` | Where to save archives (default: `backup`, relative to cwd) |
| `-cursor-user path` | Override if Cursor’s data isn’t in the usual place on your system |
| `chatsafe -version` | Print version (no `backup` needed) |

Running `chatsafe` with no arguments prints a short usage message and exits.
