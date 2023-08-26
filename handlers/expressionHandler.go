package handlers

import (
	"fmt"
	"kalvium/controllers"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Handler struct {
	Controller controllers.Controller
}

func (h *Handler) ExpressionHandler(w http.ResponseWriter, r *http.Request) {
	// Get the URL path without the leading slash
	path := mux.Vars(r)["path"]

	// Split the path into an array of values using the "/" separator
	values := strings.Split(path, "/")

	// Set header to indicate the response is a json
	w.Header().Set("Content-Type", "application/json")

	// Generated Expression
	answer, expression, err := h.Controller.ExpressionController(values)
	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err)
	}

	fmt.Fprintf(w, `{"question": "%s", "answer": "%.2f"}`, expression, answer)
}
