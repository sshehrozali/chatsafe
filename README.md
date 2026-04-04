# Chatsafe

Back up Cursor to a dated archive on disk—one command, one `.tar.gz` in a folder you choose.

**Portable:** drop the binary in a folder on your **`PATH`** and run `chatsafe` from **any directory**—no need to `cd` into a project or tool folder each time.

**Before you run:** quit Cursor so files aren’t locked.

## Install

Commands below install into **`~/.local/bin`** (macOS/Linux) or **`%USERPROFILE%\bin`** (Windows). Then add that folder to your **`PATH`** so the shell can find `chatsafe` everywhere.

### macOS (Apple Silicon, M1/M2/M3)

```bash
mkdir -p ~/.local/bin && cd ~/.local/bin
curl -fsSL -o chatsafe.tgz "https://github.com/sshehrozali/chatsafe/releases/latest/download/chatsafe-darwin-arm64.tar.gz" \
  && tar -xzf chatsafe.tgz && rm chatsafe.tgz \
  && chmod +x chatsafe
```

### macOS (Intel)

```bash
mkdir -p ~/.local/bin && cd ~/.local/bin
curl -fsSL -o chatsafe.tgz "https://github.com/sshehrozali/chatsafe/releases/latest/download/chatsafe-darwin-amd64.tar.gz" \
  && tar -xzf chatsafe.tgz && rm chatsafe.tgz \
  && chmod +x chatsafe
```

### Linux (x86_64)

```bash
mkdir -p ~/.local/bin && cd ~/.local/bin
curl -fsSL -o chatsafe.tgz "https://github.com/sshehrozali/chatsafe/releases/latest/download/chatsafe-linux-amd64.tar.gz" \
  && tar -xzf chatsafe.tgz && rm chatsafe.tgz \
  && chmod +x chatsafe
```

### Linux (arm64)

```bash
mkdir -p ~/.local/bin && cd ~/.local/bin
curl -fsSL -o chatsafe.tgz "https://github.com/sshehrozali/chatsafe/releases/latest/download/chatsafe-linux-arm64.tar.gz" \
  && tar -xzf chatsafe.tgz && rm chatsafe.tgz \
  && chmod +x chatsafe
```

### Add to PATH (macOS & Linux)

Pick the file your shell reads on startup (`~/.zshrc` for zsh, `~/.bashrc` for bash):

```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

Use `~/.bashrc` instead of `~/.zshrc` if you use bash. Open a **new** terminal, then run `chatsafe -version` from any folder.

### Windows (x64, PowerShell)

```powershell
New-Item -ItemType Directory -Force -Path "$HOME\bin" | Out-Null
Set-Location "$HOME\bin"
$base = "https://github.com/sshehrozali/chatsafe/releases/latest/download"
Invoke-WebRequest -Uri "$base/chatsafe-windows-amd64.zip" -OutFile ".\chatsafe.zip"
Expand-Archive -Path ".\chatsafe.zip" -DestinationPath "." -Force
Remove-Item ".\chatsafe.zip"
```

### Add to PATH (Windows)

User scope (persists for your account):

```powershell
$bin = [System.IO.Path]::Combine($HOME, "bin")
$old = [Environment]::GetEnvironmentVariable("Path", "User")
if ($old -notlike "*$bin*") {
  [Environment]::SetEnvironmentVariable("Path", "$old;$bin", "User")
}
```

Close and reopen PowerShell (or Windows Terminal), then run `chatsafe -version` from any directory.

Needs `curl` and `tar` on macOS/Linux.

**Other ways:** [all releases](https://github.com/sshehrozali/chatsafe/releases) · `go install github.com/sshehrozali/chatsafe/cmd/chatsafe@latest`

## Usage

```bash
chatsafe
```

Creates `backup/cursor-backup_YYYYMMDD_HHMMSS.tar.gz` in your **current working directory** (the `backup` folder is created if needed).

| Flag | What it does |
|------|----------------|
| `-out dir` | Where to save archives (default: `backup`) |
| `-cursor-user path` | Override if Cursor’s data isn’t in the usual place on your system |
| `-version` | Print version |
