package main

import (
	"fmt"
	"kalvium/configuration"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Get a new Config context
	config := configuration.NewConfig()

	router := mux.NewRouter()

	router.HandleFunc("/{history:history\\/?}", config.Handler.HistoryHandler).Methods(http.MethodGet)
	router.HandleFunc("/{path:.*\\/?}", config.Handler.ExpressionHandler).Methods(http.MethodGet)

	fmt.Println("Server listening on :8080")

	http.ListenAndServe(":8080", router)
}
