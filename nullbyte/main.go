package main

import "fmt"

func main() {
	fmt.Println("hoge\000hoge")
	h := "hoge\000hoge"
	fmt.Printf("%X\n", h)

	fmt.Printf("%X\n", "\u0000")
	fmt.Printf("%X\n", "\000")
	fmt.Printf("%X\n", "\0")
}
