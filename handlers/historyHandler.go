package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	history, err := h.Controller.HistoryController()

	// Set header to indicate the response is a json
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err)
	}
	err = json.NewEncoder(w).Encode(history)
	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err)
	}
}
