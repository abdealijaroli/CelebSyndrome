package main

import (
	"io"
	"net/http"
)

func main() {
	index := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello world")
	}

	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}
