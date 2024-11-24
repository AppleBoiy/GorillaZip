package compress

import (
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"testing"
)

func createTestFile(t *testing.T, filePath string, content string) {
	t.Helper()
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file %s: %v", filePath, err)
	}
}

func cleanupTestFile(t *testing.T, filePath string) {
	t.Helper()
	err := os.Remove(filePath)
	if err != nil {
		t.Fatalf("Failed to clean up test file %s: %v", filePath, err)
	}
}

func isValidGzipFile(t *testing.T, filePath string, expectedContent string) {
	t.Helper()
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open file %s: %v", filePath, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			t.Fatalf("Failed to close file %s: %v", filePath, err)
		}
	}(file)

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		t.Fatalf("Failed to create gzip reader for file %s: %v", filePath, err)
	}
	defer func(gzipReader *gzip.Reader) {
		err := gzipReader.Close()
		if err != nil {
			t.Fatalf("Failed to close gzip reader for file %s: %v", filePath, err)
		}
	}(gzipReader)

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gzipReader)
	if err != nil {
		t.Fatalf("Failed to read decompressed content: %v", err)
	}

	actualContent := buf.String()
	if actualContent != expectedContent {
		t.Errorf("Decompressed content mismatch. Expected: %q, Got: %q", expectedContent, actualContent)
	}
}

func TestCompressFile(t *testing.T) {
	inputFile := "test_input.txt"
	outputFile := "test_output.gz"
	testContent := "This is a sample test file for compression."

	createTestFile(t, inputFile, testContent)
	defer cleanupTestFile(t, inputFile)
	defer cleanupTestFile(t, outputFile)

	err := CompressFile(inputFile, outputFile)
	if err != nil {
		t.Fatalf("CompressFile returned an error: %v", err)
	}

	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Fatalf("Output file %s was not created", outputFile)
	}

	isValidGzipFile(t, outputFile, testContent)
}
