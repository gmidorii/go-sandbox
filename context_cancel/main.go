package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go heavy(ctx, "hoge")
	time.Sleep(3 * time.Second)
	cancel()
	fmt.Println("done cancel")

	if true {
		ctx, cancel = context.WithCancel(context.Background())
		go heavy(ctx, "fuga")
		time.Sleep(3 * time.Second)
	}
	cancel()
	fmt.Println("done cancel")
}

func heavy(ctx context.Context, s string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancel func!!")
			return
		case <-time.After(1 * time.Second):
			fmt.Println(s)
		}
	}
}
