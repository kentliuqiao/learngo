package main

import (
	"net/http"

	"github.com/kentliuqiao/learngo/crawler/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
