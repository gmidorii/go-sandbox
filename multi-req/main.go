package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/trace"

	"golang.org/x/sync/errgroup"
)

func main() {
	t, _ := os.Create("./trace.out")
	trace.Start(t)
	defer trace.Stop()

	sem := make(chan int8, 3)

	words := []string{"りんご", "バナナ", "Apple", "きつね", "ねこ", "たぬき"}
	mResult := make([]string, len(words))

	eg, ctx := errgroup.WithContext(context.Background())

	for i, w := range words {
		i, w := i, w
		eg.Go(func() error {
			sem <- 1

			body, err := request(ctx, w)
			if err != nil {
				<-sem
				return errors.New("failed")
			}

			mResult[i] = body
			<-sem
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("DONE")
}

func request(ctx context.Context, w string) (string, error) {
	v := url.Values{}
	v.Set("q", w)
	url := fmt.Sprintf("https://www.google.co.jp/search?%v", v.Encode())

	fmt.Println(url)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	rctx := r.WithContext(ctx)

	client := http.Client{}
	res, err := client.Do(rctx)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errors.New("ERROR")
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
