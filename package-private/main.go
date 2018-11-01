package main

import (
	"fmt"
	"reflect"

	"github.com/midorigreen/go-sandbox/package-private/p"
)

func main() {
	h := p.NewHoge()
	//hogeConst := h.Const()
	fmt.Println(reflect.TypeOf(hogeConst))

	h.Print()
}
