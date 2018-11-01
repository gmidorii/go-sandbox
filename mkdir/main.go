package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func createDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed create log dir: %v", err)
		}
	}
	return nil
}

func main() {
	dir := "dir"
	if err := createDir(dir); err != nil {
		log.Fatal(err)
	}

	dirDir := filepath.Join(dir, "dir")
	if err := createDir(dirDir); err != nil {
		log.Fatal(err)
	}
}
