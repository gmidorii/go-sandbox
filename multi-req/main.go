package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

func main() {
	sem := make(chan int8, 3)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	words := []string{"りんご", "バナナ", "Apple", "きつね", "ねこ", "たぬき"}
	mResult := make(chan string, len(words))
	var wg sync.WaitGroup
	for _, w := range words {
		wg.Add(1)
		go func(ctx context.Context, w string) {
			defer wg.Done()

			sem <- 1
			v := url.Values{}
			v.Set("q", w)
			url := fmt.Sprintf("https://www.google.co.jp/search?%v", v.Encode())

			result := request(ctx, url)

			select {
			case <-result:
				mResult <- w
			case <-ctx.Done():
			}
			<-sem
		}(ctx, w)
	}
	wg.Wait()

	for _, m := range <-mResult {
		fmt.Println(m)
	}
}

func request(ctx context.Context, url string) chan string {
	fmt.Println(url)
	var result chan string
	go func() {
		r, _ := http.NewRequest("GET", url, nil)
		rctx := r.WithContext(ctx)

		client := http.Client{}
		res, _ := client.Do(rctx)
		defer res.Body.Close()

		b, _ := ioutil.ReadAll(res.Body)
		result <- string(b)
	}()
	return result
}
