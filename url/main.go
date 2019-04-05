package main

import (
	"flag"
	"fmt"
	"net/url"
	"strings"
)

func main() {
	u := flag.String("u", "", "")
	flag.Parse()

	o, _ := url.Parse(*u)

	fmt.Printf("%+v\n", *o)
	fmt.Printf("%v%v\n", strings.Split(o.Host, ":")[0], o.Path)
}
