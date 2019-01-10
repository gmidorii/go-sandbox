package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var t = template.Must(template.New("static").ParseGlob("./template/*.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := t.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println(err)
	}
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	err := t.ExecuteTemplate(w, "page.html", nil)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/page", pageHandler)

	port := "5555"
	log.Printf("start port=%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
