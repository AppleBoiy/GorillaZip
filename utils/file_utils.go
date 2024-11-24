package utils

import (
	"os"
	"path/filepath"
)

func ListFiles(dir string) ([]string, error) {
	//- Traverse the directory recursively.
	//- Skip non-file entries (e.g., directories).

	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
