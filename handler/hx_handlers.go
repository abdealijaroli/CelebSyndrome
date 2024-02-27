package handler

import (
	"fmt"
	"net/http"

	"github.com/abdealijaroli/gostream/util"
)

func CelebHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
	}

	celeb := r.PostFormValue("celeb")
	if celeb != "" {
		fmt.Println("Celeb:", celeb)
		resp := util.GenerateAIResponse(celeb)
		fmt.Fprint(w, resp)
	} else {
		fmt.Fprintf(w, "No name provided!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
