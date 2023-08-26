package handlers

import (
	"encoding/json"
	"kalvium/controllers"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Handler struct {
	Controller controllers.Controller
}

type errorResponse struct {
	Error error `json:"error"`
}

type expressionResponse struct {
	Question string  `json:"question"`
	Answer   float64 `json:"answer"`
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
		resp := errorResponse{
			Error: err,
		}
		err = json.NewEncoder(w).Encode(resp)
		return
	}

	resp := expressionResponse{
		Question: expression,
		Answer:   answer,
	}
	err = json.NewEncoder(w).Encode(resp)
}
