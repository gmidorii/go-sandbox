package main

import "fmt"

type Iota int

const (
	_ Iota = iota
	One
	Two
	Three
	Four
	Five
)

func main() {
	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Three)
	fmt.Println(Four)
	fmt.Println(Five)
}
