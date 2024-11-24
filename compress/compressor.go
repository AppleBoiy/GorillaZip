package compress

import (
	"compress/gzip"
	"io"
	"os"
)

func CompressFile(input, output string) error {
	//- Open the input file.
	//- Create a new gzip writer.
	//- Copy the input file to the gzip writer.
	//- Close the gzip writer.
	//- Close the input file.

	inputFile, err := os.Open(input)
	if err != nil {
		return err
	}
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			panic(err)
		}
	}(inputFile)

	outputFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			panic(err)
		}
	}(outputFile)

	gzipWriter := gzip.NewWriter(outputFile)
	defer func(gzipWriter *gzip.Writer) {
		err := gzipWriter.Close()
		if err != nil {
			panic(err)
		}
	}(gzipWriter)

	_, err = io.Copy(gzipWriter, inputFile)
	if err != nil {
		return err
	}

	return nil
}
