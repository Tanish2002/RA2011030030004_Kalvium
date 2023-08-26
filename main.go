package main

import (
	"fmt"
	"kalvium/configuration"
	"net/http"
)

func main() {

	// Get a new Config context
	config := configuration.NewConfig()

	http.HandleFunc("/", config.Handler.ExpressionHandler)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
