package main

import (
	"runtime"
	"sync"
	"fmt"
)

func forever(i int) {
	fmt.Println(i)
	for {
		runtime.Gosched()
	}
}

func main()  {
	fmt.Printf("Go Maxprocs: %v\n", runtime.GOMAXPROCS(-1))
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go forever(i)
	}
	wg.Wait()
}