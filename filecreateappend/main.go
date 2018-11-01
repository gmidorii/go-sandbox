package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.OpenFile("hoge.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	defer f.Close()

	fmt.Fprintf(f, "hoge\n")
	f.Close()

	f2, _ := os.OpenFile("hoge.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	defer f2.Close()

	fmt.Fprintf(f2, "hoge2\n")

}
