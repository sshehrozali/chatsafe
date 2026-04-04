// Chatsafe backs up Cursor IDE user data (workspaceStorage, settings, etc.) into a single timestamped .tar.gz.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/sshehrozali/chatsafe/internal/pack"
)

// version is set at link time by GoReleaser (-X main.version=...).
var version = "dev"

func defaultCursorUser() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	switch runtime.GOOS {
	case "darwin":
		return filepath.Join(home, "Library", "Application Support", "Cursor", "User")
	case "windows":
		if appData := os.Getenv("APPDATA"); appData != "" {
			return filepath.Join(appData, "Cursor", "User")
		}
		return ""
	default:
		// Linux and other Unix-like systems
		if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
			return filepath.Join(xdg, "Cursor", "User")
		}
		return filepath.Join(home, ".config", "Cursor", "User")
	}
}

func main() {
	log.SetFlags(0)

	out := flag.String("out", "backup", "output directory for archives; created on first run, reused afterward (relative to cwd)")
	cursorUser := flag.String("cursor-user", "", "path to Cursor User directory (default: standard location for this OS)")
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	root := *cursorUser
	if root == "" {
		root = defaultCursorUser()
	}
	if root == "" {
		log.Fatal("chatsafe: could not resolve default Cursor User path")
	}

	stopSpinner := startSpinner("Creating backup…")
	archivePath, err := pack.Pack(root, *out)
	stopSpinner()
	if err != nil {
		log.Fatalf("chatsafe: %v", err)
	}

	fi, err := os.Stat(archivePath)
	if err != nil {
		log.Fatalf("chatsafe: stat archive: %v", err)
	}

	fmt.Printf("backup written: %s (%d bytes)\n", archivePath, fi.Size())
}

// stderrIsTTY reports whether stderr is an interactive terminal (skip spinner when piped).
func stderrIsTTY() bool {
	fi, err := os.Stderr.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

// startSpinner draws a Braille animation on stderr until stop is called.
func startSpinner(message string) (stop func()) {
	if !stderrIsTTY() {
		return func() {}
	}

	done := make(chan struct{})
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

	go func() {
		tick := time.NewTicker(80 * time.Millisecond)
		defer tick.Stop()
		i := 0
		for {
			select {
			case <-done:
				_, _ = fmt.Fprint(os.Stderr, "\r\033[K")
				return
			case <-tick.C:
				_, _ = fmt.Fprintf(os.Stderr, "\r%s %s", frames[i%len(frames)], message)
				i++
			}
		}
	}()

	return func() { close(done) }
}
