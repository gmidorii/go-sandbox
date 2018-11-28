package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/sh", "-c", "cd ./dir && ls")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
