package main

import (
	"fmt"
	"time"
)

func main() {

	s1 := time.Now()
	sa1 := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		sa1 = append(sa1, i)
	}
	measure("len + append", s1, sa1)

	s2 := time.Now()
	sa2 := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		sa2 = append(sa2, i)
	}
	measure("cap + append", s2, sa2)

	var sa3 []int
	s3 := time.Now()
	for i := 0; i < 100000; i++ {
		sa3 = append(sa3, i)
	}
	measure("don not initialize + append", s3, sa3)

	s4 := time.Now()
	sa4 := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		sa4[i] = i
	}
	measure("len + index", s4, sa4)

}

func measure(title string, t time.Time, s []int) {
	fmt.Println("---------------------------")
	fmt.Printf("%v: \n%v nsec\n", title, time.Now().Sub(t).Nanoseconds())
	fmt.Printf("len %v, 1: %v \n", len(s), s[1])
	fmt.Println("---------------------------")
}
