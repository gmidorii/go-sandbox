package main

import (
	"log"
	"time"
)

func heavy() {
	time.Sleep(3 * time.Second)
}

func measureTime(message string) func() {
	log.Printf("===== START %v =====", message)
	start := time.Now()
	return func() {
		log.Printf("===== END %v %v sec =====", message, time.Now().Sub(start))
	}
}

func main() {
	end := measureTime("heavy")
	heavy()
	end()
}
