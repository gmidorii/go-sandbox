package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/twpayne/go-geom/encoding/ewkb"
	"github.com/twpayne/go-geom/encoding/wkt"
)

func main() {
	// EWKB
	s := ""
	fmt.Println(len(s))

	r, _ := hex.DecodeString(s)
	fmt.Println(r)
	g, err := ewkb.Unmarshal(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g.SRID())
	fmt.Println(g.Layout())
	fmt.Println(g.FlatCoords())

	w, _ := wkt.Marshal(g)
	fmt.Println(w)
}
