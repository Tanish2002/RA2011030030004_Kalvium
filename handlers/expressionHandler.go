package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type expressionResponse struct {
	Question string  `json:"question"`
	Answer   float64 `json:"answer"`
}

func (h *Handler) ExpressionHandler(w http.ResponseWriter, r *http.Request) {
	// Get the URL path without the leading slash
	path := mux.Vars(r)["path"]

	// Split the path into an array of values using the "/" separator
	values := strings.Split(path, "/")

	// Generated Expression
	answer, expression, err := h.Controller.ExpressionController(values)
	if ok := h.errorResp(err, w); ok {
		return
	}

	resp := expressionResponse{
		Question: expression,
		Answer:   answer,
	}

	// Set header to indicate the response is a json
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if ok := h.errorResp(err, w); ok {
		return
	}
}
