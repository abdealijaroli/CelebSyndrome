package handler

import (
	"fmt"
	"net/http"
)

func CelebHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		fmt.Println("This is an htmx request")
	}

	fmt.Println("Request method:", r.Method)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	fmt.Println("Form data:", r.Form)

	celeb := r.PostFormValue("celeb")
	if celeb == "" {
		fmt.Fprintf(w, "No name provided!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("Celeb:", celeb)
}
