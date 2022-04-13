package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var fileExtension string
	fileType := flag.String("f", "", "sratch file type")
	fileName := flag.String("n", "tmp", "sratch file name")
	flag.Parse()

	switch *fileType {
	case "go", "golang":
		fileExtension = "go"
	case "python", "py":
		fileExtension = "py"
	case "bash", "shell", "sh":
		fileExtension = "sh"
	}

	file := fmt.Sprintf("%v.%v", *fileName, fileExtension)
	os.Create(file)
}
