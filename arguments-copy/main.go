package main

import "fmt"

type Hoge struct {
	Name string
}

func main() {
	h := Hoge{Name: "Default"}

	// not overwrite
	overwrite(h)
	fmt.Println(h.Name)

	// overwrite
	overwritePointer(&h)
	fmt.Println(h.Name)
}

func overwrite(h Hoge) {
	h.Name = "Overwrite"
}

func overwritePointer(h *Hoge) {
	h.Name = "OverwritePointer"
}
