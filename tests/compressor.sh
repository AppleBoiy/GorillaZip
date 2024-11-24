#!/bin/bash

#---
# Use for test compressor
#---

# Set up environment variables (if needed)
# export GOPATH=$HOME/go
# export GOROOT=/usr/local/go
# export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

PROJECT_DIR="."
cd "$PROJECT_DIR" || { echo "Project directory not found!"; exit 1; }

echo "Cleaning up previous builds..."
rm -f gorilla test_output.gz test_input.txt

echo "Creating test input file..."
echo "This is a test file for compression." > test_input.txt

echo "Building the Go project..."
go build -o gorilla || { echo "Go build failed!"; exit 1; }

echo "Running the compression tool..."
./gorilla test_input.txt test_output.gz || { echo "Compression failed!"; exit 1; }

if [ -f "test_output.gz" ]; then
    echo "Compressed file test_output.gz created successfully."
else
    echo "Compressed file test_output.gz was not created!"
    exit 1
fi

echo "Verifying the content of the compressed file..."
decompressed_content=$(gunzip -c test_output.gz)

expected_content="This is a test file for compression."
if [ "$decompressed_content" == "$expected_content" ]; then
    echo "Decompression successful: Content matches expected."
else
    echo "Decompression failed: Content does not match!"
    exit 1
fi

echo "Cleaning up test files..."
rm -f test_input.txt test_output.gz gorilla

echo "Manual testing completed successfully!"
