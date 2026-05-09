// Tests for llmsave pack archives (no real Cursor install required).
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPack_CreatesArchiveWithExpectedTopLevel(t *testing.T) {
	cursorRoot := t.TempDir()
	outDir := t.TempDir()

	if err := os.MkdirAll(filepath.Join(cursorRoot, "workspaceStorage", "ws1"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(cursorRoot, "workspaceStorage", "ws1", "note.txt"), []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(cursorRoot, "settings.json"), []byte("{}"), 0o644); err != nil {
		t.Fatal(err)
	}

	archivePath, err := Pack(cursorRoot, outDir)
	if err != nil {
		t.Fatalf("Pack: %v", err)
	}

	matches, err := filepath.Glob(filepath.Join(outDir, "llmsave-backup_*.tar.gz"))
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 1 {
		t.Fatalf("expected exactly one backup archive, got %d: %v", len(matches), matches)
	}
	if matches[0] != archivePath {
		t.Fatalf("glob path %q != returned path %q", matches[0], archivePath)
	}

	top, err := ListTopLevelNames(archivePath)
	if err != nil {
		t.Fatalf("ListTopLevelNames: %v", err)
	}
	for _, want := range []string{"workspaceStorage", "settings.json"} {
		if _, ok := top[want]; !ok {
			t.Fatalf("missing top-level entry %q; got %v", want, top)
		}
	}
}

func TestPack_ErrWhenNothingToArchive(t *testing.T) {
	emptyRoot := t.TempDir()
	outDir := t.TempDir()
	_, err := Pack(emptyRoot, outDir)
	if err == nil {
		t.Fatal("expected error when Cursor user tree is empty")
	}
}
