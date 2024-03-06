package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/abdealijaroli/gostream/handler"
	"github.com/abdealijaroli/gostream/util"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	db, err := util.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	err = util.CreateUserTable(db)
	if err != nil {
		log.Fatal("Error creating user table:", err)
	}

	fs := http.FileServer(http.Dir("./view"))
	http.Handle("/", fs)

	http.HandleFunc("/celeb", handler.CelebHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server is running on port", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}