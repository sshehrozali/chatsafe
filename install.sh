#!/bin/sh
# One-shot install: download latest chatsafe, put it in ~/.local/bin, add that dir to PATH in your shell rc.
set -e

BIN_DIR="${BIN_DIR:-$HOME/.local/bin}"
mkdir -p "$BIN_DIR"

case $(uname -s) in
Darwin) os=darwin ;;
Linux) os=linux ;;
*)
	echo "install.sh: use PowerShell on Windows:" >&2
	echo '  irm https://raw.githubusercontent.com/sshehrozali/chatsafe/main/install.ps1 | iex' >&2
	exit 1
	;;
esac

case $(uname -m) in
x86_64 | amd64) arch=amd64 ;;
arm64 | aarch64) arch=arm64 ;;
*)
	echo "install.sh: unsupported CPU ($(uname -m))." >&2
	exit 1
	;;
esac

OUT="$BIN_DIR/chatsafe.tgz"

# Resolve the real asset URL from the API (avoids 404 when stable-named files are missing from a release).
echo "Fetching latest release for ${os}/${arch} ..."
GH_API="https://api.github.com/repos/sshehrozali/chatsafe/releases/latest"
if [ -n "${GITHUB_TOKEN:-}" ]; then
	JSON=$(curl -fsSL \
		-H "Accept: application/vnd.github+json" \
		-H "User-Agent: chatsafe-install" \
		-H "Authorization: Bearer $GITHUB_TOKEN" \
		"$GH_API")
else
	JSON=$(curl -fsSL \
		-H "Accept: application/vnd.github+json" \
		-H "User-Agent: chatsafe-install" \
		"$GH_API")
fi

# Prefer stable asset name, else versioned chatsafe_X.Y.Z_os_arch.tar.gz
URL=$(printf '%s' "$JSON" | tr '"' '\n' | grep -E '^https://github\.com/.*chatsafe-'"${os}"'-'"${arch}"'\.tar\.gz$' | head -n 1)
if [ -z "$URL" ]; then
	URL=$(printf '%s' "$JSON" | tr '"' '\n' | grep -E '^https://github\.com/.*chatsafe_[^_]+_'"${os}"'_'"${arch}"'\.tar\.gz$' | head -n 1)
fi
if [ -z "$URL" ]; then
	echo "install.sh: no ${os}/${arch} .tar.gz in latest release." >&2
	echo "See https://github.com/sshehrozali/chatsafe/releases" >&2
	rm -f "$OUT"
	exit 1
fi

echo "Downloading ..."
curl -fsSL "$URL" -o "$OUT"

tar -xzf "$BIN_DIR/chatsafe.tgz" -C "$BIN_DIR"
rm -f "$BIN_DIR/chatsafe.tgz"
chmod +x "$BIN_DIR/chatsafe"

if [ -f "$HOME/.zshrc" ]; then
	rc="$HOME/.zshrc"
elif [ -f "$HOME/.bashrc" ]; then
	rc="$HOME/.bashrc"
else
	rc="$HOME/.profile"
	touch "$rc"
fi

if ! grep -q '# chatsafe PATH' "$rc" 2>/dev/null; then
	{
		echo ""
		echo "# chatsafe PATH"
		echo "export PATH=\"$BIN_DIR:\$PATH\""
	} >>"$rc"
	echo "Added $BIN_DIR to PATH in $rc"
else
	echo "PATH already configured for chatsafe in $rc (skipped duplicate line)"
fi

echo ""
echo "Done."
echo ""
echo "This terminal still has the old PATH. Do one of:"
echo "  1) source $rc && chatsafe -version"
echo "  2) Open a new terminal, then: chatsafe -version"
echo "  3) Run now without PATH: $BIN_DIR/chatsafe -version"
echo ""
echo "Create a backup: chatsafe backup"
