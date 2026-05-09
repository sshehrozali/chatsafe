# Chatsafe

**Back up your LLM memory**—chats, workspace storage, and IDE state—into a single dated `.tar.gz`. One command, predictable archives, full control over where they land.

Use it to **protect context** before an OS reinstall, **snapshot** a project era, **move machines**, or keep an **offline copy** of what your assistant “remembers” on disk.

**Portable:** install once with the command below; **`chatsafe`** lands on your **`PATH`** so you can run it from **any directory**.

**Supported IDEs**

- **Cursor** — supported
- **Claude** — coming soon

**Before you run:** quit your **IDE** so databases and files are not locked mid-backup.

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

Writes `backup/cursor-backup_YYYYMMDD_HHMMSS.tar.gz` under your **current working directory** (the `backup` folder is created if needed). That archive holds the **LLM IDE user data tree** Chatsafe packages from your machine (naming stays `cursor-backup_*` for compatibility).

### Examples

```bash
# Default output directory under your current directory
cd ~
chatsafe backup
```

```bash
# Save archives somewhere fixed (folder is created if needed)
chatsafe backup -out ~/Documents/llm-backups
```

```bash
# LLM IDE data lives somewhere non-standard—point at the User folder explicitly
chatsafe backup -cursor-user "/path/to/Cursor/User"
```

| Command / flag | What it does |
|----------------|--------------|
| `chatsafe backup` | Create a timestamped archive of LLM IDE user data |
| `-out dir` | Where to save archives (default: `backup`, relative to cwd) |
| `-cursor-user path` | Override the **Cursor** `User` directory if yours is not in the usual OS location |
| `chatsafe -version` | Print version (no `backup` needed) |

Running `chatsafe` with no arguments prints a short usage message and exits.
