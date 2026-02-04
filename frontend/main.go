package main

import (
	"ClubManager/frontend/internal/handler"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HelloHandler)
  mux.HandleFunc("/user", handler.UserHandler)

  fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mux)
}
