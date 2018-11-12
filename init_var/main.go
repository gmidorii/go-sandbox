package main

import "fmt"

func initVar(a *int, s *string) {
	ai := -1
	si := "hoge"
	a = &ai
	s = &si
}

func initVar2() (a int, s string) {
	a = -1
	s = "hoge"

	return
}

func main() {
	var a int
	var s string
	fmt.Println(a)
	fmt.Println(s)

	initVar(&a, &s)
	fmt.Println(a)
	fmt.Println(s)

	a, s = initVar2()
	fmt.Println(a)
	fmt.Println(s)
}
