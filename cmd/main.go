package main

import (
	"net/http"

	"github.com/abdealijaroli/gostream/handler"
)

func main() {
	fs := http.FileServer(http.Dir("./view"))
	http.Handle("/", fs)

	http.HandleFunc("/celeb", handler.CelebHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
