// Package pack writes llmsave gzip-tar archives of Cursor's User directory (see llmsave/cmd/llmsave).
package pack

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// CursorUserEntries are top-level names under Cursor's User directory, aligned with bin/quick_backup.sh.
var CursorUserEntries = []string{
	"workspaceStorage",
	"globalStorage",
	"History",
	"settings.json",
	"keybindings.json",
}

// Pack writes a llmsave gzip-compressed tar of the given entries under cursorUserRoot into outDir.
// outDir is created with MkdirAll if missing; later runs can reuse the same directory for more archives.
// The archive is named llmsave-backup_YYYYMMDD_HHMMSS.tar.gz.
// Missing paths are skipped. Returns an error if nothing was archived.
func Pack(cursorUserRoot, outDir string) (archivePath string, err error) {
	cursorUserRoot = filepath.Clean(cursorUserRoot)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return "", err
	}

	ts := time.Now().Format("20060102_150405")
	baseName := fmt.Sprintf("llmsave-backup_%s.tar.gz", ts)
	archivePath = filepath.Join(outDir, baseName)

	f, err := os.Create(archivePath)
	if err != nil {
		return "", err
	}
	defer func() {
		if err != nil {
			_ = f.Close()
			_ = os.Remove(archivePath)
		}
	}()

	gzw := gzip.NewWriter(f)
	tw := tar.NewWriter(gzw)

	var count int
	for _, ent := range CursorUserEntries {
		abs := filepath.Join(cursorUserRoot, ent)
		st, statErr := os.Stat(abs)
		if os.IsNotExist(statErr) {
			continue
		}
		if statErr != nil {
			err = statErr
			return "", err
		}
		if st.IsDir() {
			n, walkErr := addDir(tw, ent, abs)
			if walkErr != nil {
				err = walkErr
				return "", err
			}
			count += n // includes directory headers and files
		} else {
			if addErr := addFile(tw, ent, abs, st); addErr != nil {
				err = addErr
				return "", err
			}
			count++
		}
	}

	if count == 0 {
		err = fmt.Errorf("no Cursor user data found under %s", cursorUserRoot)
		return "", err
	}

	if err = tw.Close(); err != nil {
		return "", err
	}
	if err = gzw.Close(); err != nil {
		return "", err
	}
	if err = f.Close(); err != nil {
		return "", err
	}
	err = nil
	return archivePath, nil
}

func addDir(tw *tar.Writer, arcPrefix, absDir string) (int, error) {
	var count int
	err := filepath.WalkDir(absDir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, err := filepath.Rel(absDir, path)
		if err != nil {
			return err
		}
		arcName := arcPrefix
		if rel != "." {
			arcName = filepath.Join(arcPrefix, rel)
		}
		arcName = filepath.ToSlash(arcName)

		info, err := d.Info()
		if err != nil {
			return err
		}
		if d.IsDir() {
			hdr, err := tar.FileInfoHeader(info, "")
			if err != nil {
				return err
			}
			hdr.Name = arcName + "/"
			if err := tw.WriteHeader(hdr); err != nil {
				return err
			}
			count++
			return nil
		}
		if err := addFile(tw, arcName, path, info); err != nil {
			return err
		}
		count++
		return nil
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func addFile(tw *tar.Writer, arcName, absPath string, info fs.FileInfo) error {
	if info.Mode()&os.ModeSymlink != 0 {
		return nil
	}
	if !info.Mode().IsRegular() {
		return nil
	}

	hdr, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}
	hdr.Name = filepath.ToSlash(arcName)
	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}

	r, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer r.Close()

	if _, err := io.Copy(tw, r); err != nil {
		return err
	}
	return nil
}

// ListTopLevelNames returns the set of top-level path names inside a llmsave-backup_*.tar.gz file.
func ListTopLevelNames(archivePath string) (map[string]struct{}, error) {
	f, err := os.Open(archivePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	out := make(map[string]struct{})
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		name := strings.TrimSuffix(hdr.Name, "/")
		if name == "" {
			continue
		}
		first := strings.SplitN(name, "/", 2)[0]
		out[first] = struct{}{}
		if _, err := io.Copy(io.Discard, tr); err != nil {
			return nil, err
		}
	}
	return out, nil
}
