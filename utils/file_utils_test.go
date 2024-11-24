package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func setupTestDirectory(t *testing.T, baseDir string) {
	t.Helper()
	err := os.MkdirAll(filepath.Join(baseDir, "subdir"), 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	err = os.MkdirAll(filepath.Join(baseDir, "subdir", "subsubdir"), 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	files := []string{
		"file1.txt",
		"file2.txt",
		"subdir/file3.txt",
		"subdir/.hidden",
		"subdir/subsubdir/file4.txt",
	}
	for _, file := range files {
		filePath := filepath.Join(baseDir, file)
		err := os.WriteFile(filePath, []byte("sample content"), 0644)
		if err != nil {
			t.Fatalf("Failed to create file %s: %v", filePath, err)
		}
	}
}

func teardownTestDirectory(t *testing.T, baseDir string) {
	t.Helper()
	err := os.RemoveAll(baseDir)
	if err != nil {
		t.Fatalf("Failed to clean up test directory: %v", err)
	}
}

func TestListFiles(t *testing.T) {
	baseDir := "testdata"
	setupTestDirectory(t, baseDir)
	defer teardownTestDirectory(t, baseDir)

	// Expected files
	expectedFiles := map[string]bool{
		filepath.Join(baseDir, "file1.txt"):                  true,
		filepath.Join(baseDir, "file2.txt"):                  true,
		filepath.Join(baseDir, "subdir/file3.txt"):           true,
		filepath.Join(baseDir, "subdir/.hidden"):             true,
		filepath.Join(baseDir, "subdir/subsubdir/file4.txt"): true,
	}

	files, err := ListFiles(baseDir)
	if err != nil {
		t.Fatalf("ListFiles returned an error: %v", err)
	}

	if len(files) != len(expectedFiles) {
		t.Errorf("Expected %d files, but got %d", len(expectedFiles), len(files))
	}

	for _, file := range files {
		if !expectedFiles[file] {
			t.Errorf("Unexpected file: %s", file)
		}
		delete(expectedFiles, file) // Mark as found
	}

	if len(expectedFiles) > 0 {
		t.Errorf("Some expected files were not found: %v", expectedFiles)
	}
}
