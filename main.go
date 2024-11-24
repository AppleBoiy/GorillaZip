package main

import (
	fileutils "GorillaZip/utils"
)

func main() {
	files, err := fileutils.ListFiles("./.git")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		println(file)
	}
}
