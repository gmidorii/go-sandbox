package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	sem := make(chan int8, 3)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	words := []string{"りんご", "バナナ", "Apple", "きつね", "ねこ", "たぬき"}
	mResult := make(chan string, len(words))
	for _, w := range words {
		go func(ctx context.Context, w string) {
			sem <- 1
			v := url.Values{}
			v.Set("q", w)
			url := fmt.Sprintf("https://www.google.co.jp/search?%v", v.Encode())

			request(ctx, url)
			mResult <- w
			<-sem
		}(ctx, w)
	}

	for m := range mResult {
		fmt.Println(m)
	}
}

func request(ctx context.Context, url string) string {
	fmt.Println(url)
	r, _ := http.NewRequest("GET", url, nil)
	rctx := r.WithContext(ctx)

	client := http.Client{}
	res, _ := client.Do(rctx)
	defer res.Body.Close()

	b, _ := ioutil.ReadAll(res.Body)
	return string(b)
}
