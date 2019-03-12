package main

import (
	"fmt"
	"os"
)

//getFilePerm will get file permissions
func getFilePerm(fpath string) {

	fileInfo, err := os.Stat(fpath)

	// check if there is an error
	if err != nil {

		// check if error is file does not exist
		if os.IsNotExist(err) {
			fmt.Println("File does not exist.")
		}

	}

	mode := fileInfo.Mode()

	fmt.Println(fileInfo, "mode is ", mode)
}

func main() {
	fmt.Println("gvs")
}
