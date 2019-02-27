package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := flag.String("d", "", "sim directory")
	flag.Parse()

	filepath.Walk(*dir, func(path string, fileInfo os.FileInfo, err error) error {
		fmt.Println(path)
		fmt.Println(fileInfo)
		return nil
	})
}
