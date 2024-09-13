package main

import (
	"fmt"
	"go-server/handlers"
	"log"
	"net/http"
)





func main() {
	http.HandleFunc("/posts", handlers.PostsHandler)
	http.HandleFunc("/post/", handlers.PostHandler)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

