package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	index := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello world")
	}

	http.HandleFunc("/", index)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server is running on port 8080")
	}
}
