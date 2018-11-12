package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/machinebox/progress"
)

func main() {
	ctx := context.Background()
	o, _ := os.Open("./input.csv")
	defer o.Close()
	s, _ := o.Stat()
	size := s.Size()

	r := progress.NewReader(o)

	go func() {
		pChan := progress.NewTicker(ctx, r, int64(size), 10*time.Millisecond)
		for p := range pChan {
			fmt.Printf("\r %v ...", p.Remaining().Round(time.Millisecond))
		}
		fmt.Println("completed")
	}()

	time.Sleep(1 * time.Second)

	if _, err := io.Copy(ioutil.Discard, r); err != nil {
		log.Fatalln(err)
	}
}
