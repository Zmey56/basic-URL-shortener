package main

import (
	"fmt"
	"github.com/Zmey56/basic-URL-shortener/shortener"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/encode", shortener.Encode)
	http.HandleFunc("/decode/", shortener.Decode)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

//2. Error Handling: Make sure to implement proper error handling
