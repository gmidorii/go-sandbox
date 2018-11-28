package main

import "fmt"

func main() {
	backoff := []int{10, 30, 50, 80}
	for _, b := range backoff {
		backoff = append(backoff, -b)
	}

	fmt.Println(backoff)
}
