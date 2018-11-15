package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go heavy(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	fmt.Println("done cancel")
	time.Sleep(3 * time.Second)
}

func heavy(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancel func!!")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("hoge")
		}
	}
}
