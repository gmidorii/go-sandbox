package pool

import (
	"strconv"
	"sync"
)

const number = 1000000

var s []string = []string{"1", "hoge", "10000", "fuga"}

func NotConcurrent() []bool {
	result := make([]bool, number)
	for i := 0; i < number; i++ {
		result[i] = check(s)
	}
	return result
}

func Goroutine() []bool {
	type value struct {
		idx    int
		result bool
	}
	rChan := make(chan value, number)

	result := make([]bool, number)
	go func() {
		var wg sync.WaitGroup
		for i := 0; i < number; i++ {
			i := i
			wg.Add(1)
			go func() {
				defer wg.Done()
				r := check(s)
				rChan <- value{idx: i, result: r}
			}()
		}
		wg.Wait()
		close(rChan)
	}()

	for {
		v, ok := <-rChan
		if !ok {
			break
		}
		result[v.idx] = v.result
	}
	return result
}

func check(s []string) bool {
	i, err := strconv.Atoi(s[0])
	if err != nil {
		return false
	}
	if i < 0 {
		return false
	}

	if s[1] != "hoge" {
		return false
	}

	n, err := strconv.Atoi(s[0])
	if err != nil {
		return false
	}
	if n < 10000 {
		return false
	}

	return true
}
