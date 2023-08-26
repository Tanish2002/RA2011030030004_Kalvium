package handlers

import (
	"fmt"
	"kalvium/controllers"
	"net/http"
	"strings"
)

type Handler struct {
	Controller controllers.Controller
}

func (h *Handler) ExpressionHandler(w http.ResponseWriter, r *http.Request) {

	// Get the URL path without the leading slash
	path := r.URL.Path[1:]

	// Split the path into an array of values using the "/" separator
	values := strings.Split(path, "/")

	// Set header to indicate the response is a json
	w.Header().Set("Content-Type", "application/json")

	// Generated Expression
	expression, err := h.Controller.ExpressionController(values)
	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err)
	}

	fmt.Fprintf(w, `{"value": "%.2f"}`, expression)

}
