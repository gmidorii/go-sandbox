package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type request struct {
	perSecond       int
	concurrentLimit int
}

func NewRequest(perSecond, concurrentLimit int) request {
	return request{
		perSecond:       perSecond,
		concurrentLimit: concurrentLimit,
	}
}

func (r *request) request(qChan chan query, res chan response) {
	betwennMill := 1000 / r.perSecond
	tick := time.NewTicker(time.Duration(betwennMill) * time.Millisecond)
	defer tick.Stop()
	throttle := make(chan time.Time, 1)
	go func() {
		for t := range tick.C {
			throttle <- t
		}
	}()

	var wg sync.WaitGroup
	limiter := make(chan struct{}, r.concurrentLimit)
	for {
		q, ok := <-qChan
		if !ok {
			break
		}

		<-throttle
		limiter <- struct{}{}
		wg.Add(1)

		go func(q query, res chan response) {
			defer func() {
				fmt.Println("Done")
				<-limiter
				wg.Done()
			}()
			// unit request
			fmt.Println(q.rawQuery)
			body, err := unitReq(q)
			if err != nil {
				res <- response{err: err}
				return
			}
			res <- response{body: body}
		}(q, res)

	}

	wg.Wait()
	for {
		if len(res) == 0 {
			break
		}
	}
	close(res)
}

func unitReq(q query) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("http://example.com?%v", q.rawQuery))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type query struct {
	rawQuery string
}

type response struct {
	body []byte
	err  error
}

func main() {
	// generate query
	q := make(chan query, 1)
	go func() {
		// sample
		for i := 0; i < 3; i++ {
			q <- query{rawQuery: fmt.Sprintf("idx=%v", i)}
		}
		close(q)
	}()

	// request
	req := NewRequest(1, 10)
	res := make(chan response, 10)
	go req.request(q, res)

	// response
finished:
	for {
		select {
		case r, ok := <-res:
			// channel close
			if !ok {
				break finished
			}
			if r.err != nil {
				// error handling
				continue
			}

			fmt.Println(string(r.body))
		}
	}
}
