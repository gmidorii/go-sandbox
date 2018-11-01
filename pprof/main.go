package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	log.Println(http.ListenAndServe(":6060", nil))
}
