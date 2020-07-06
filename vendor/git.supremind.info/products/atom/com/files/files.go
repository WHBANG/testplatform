package files

import (
	"bufio"
	"context"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type LocalFile struct {
	Relative string
	Local    string
	FileInfo os.FileInfo
}

// returned files are basename for single file input,
// or relative filepath for directory input
func Walkthrough(ctx context.Context, root string, files chan<- *LocalFile) error {
	info, err := os.Stat(root)
	if err != nil {
		return errors.Wrap(err, "stat root file failed")
	}

	// for single file input
	if !info.IsDir() {
		files <- &LocalFile{Relative: filepath.Base(root), Local: root, FileInfo: info}
		return nil
	}

	// for directory input
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "error walking file")
		}

		if info.Mode().IsRegular() {
			rel, err := filepath.Rel(root, path)
			if err != nil {
				// shall never get here
				return errors.Wrap(err, "can not get relative path")
			}

			select {
			case files <- &LocalFile{Relative: rel, Local: path, FileInfo: info}:
			case <-ctx.Done():
				return errors.Wrap(ctx.Err(), "walkthrough interrupted")
			}
		}

		return nil
	})
}

func ReadDumppedList(ctx context.Context, listFile string) (<-chan string, <-chan error) {
	files := make(chan string)
	errs := make(chan error)

	go func() {
		defer close(errs)
		defer close(files)

		f, err := os.Open(listFile)
		if err != nil {
			errs <- errors.Wrap(err, "open list file failed")
			return
		}
		defer f.Close()

		s := bufio.NewScanner(f)
	scan:
		for s.Scan() {
			select {
			case <-ctx.Done():
				errs <- errors.New("interrupted")
				break scan
			case files <- s.Text():
			}
		}

		if s.Err() != nil {
			errs <- errors.Wrap(err, "error in scanning input list put")
		}
	}()

	return files, errs
}

// CleanPath makes a path safe for use with filepath.Join. This is done by not
// only cleaning the path, but also (if the path is relative) adding a leading
// '/' and cleaning it (then removing the leading '/'). This ensures that a
// path resulting from prepending another path will always resolve to lexically
// be a subdirectory of the prefixed path. This is all done lexically, so paths
// that include symlinks won't be safe as a result of using CleanPath.
//
// This function comes from runC (libcontainer/utils/utils.go).
func CleanPath(path string) string {
	// Deal with empty strings nicely.
	if path == "" {
		return ""
	}

	// Ensure that all paths are cleaned (especially problematic ones like
	// "/../../../../../" which can cause lots of issues).
	path = filepath.Clean(path)

	// If the path isn't absolute, we need to do more processing to fix paths
	// such as "../../../../<etc>/some/path". We also shouldn't convert absolute
	// paths to relative ones.
	if !filepath.IsAbs(path) {
		path = filepath.Clean(string(os.PathSeparator) + path)
		// This can't fail, as (by definition) all paths are relative to root.
		path, _ = filepath.Rel(string(os.PathSeparator), path)
	}

	// Clean the path again for good measure.
	return filepath.Clean(path)
}
