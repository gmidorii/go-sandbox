package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("hello"))
	x := h.Sum([]byte{0x05})
	fmt.Println(string(x))
	fmt.Println(len(x))
	fmt.Printf("%x\n", x)
}
