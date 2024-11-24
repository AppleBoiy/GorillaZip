package main

import (
	compressor "GorillaZip/compress"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <input file> <output file>", os.Args[0])
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	err := compressor.CompressFile(inputFile, outputFile)
	if err != nil {
		log.Fatalf("Error compressing file: %v", err)
	}

	fmt.Printf("File '%s' compressed successfully to '%s'\n", inputFile, outputFile)
}
