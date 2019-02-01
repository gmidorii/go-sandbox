package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		j := i
		fmt.Printf("top: %v\n", j)
		go func() {
			defer func() {
				fmt.Printf("defer: %v\n", j)
			}()
			fmt.Printf("func: %v\n", j)
		}()
		i++
	}
	time.Sleep(3 * time.Second)
}
