package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	history, err := h.Controller.HistoryController()

	// Set header to indicate the response is a json
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp := errorResponse{
			Error: err.Error(),
		}
		err = json.NewEncoder(w).Encode(resp)
		return
	}
	err = json.NewEncoder(w).Encode(history)
	if err != nil {
		resp := errorResponse{
			Error: err.Error(),
		}
		err = json.NewEncoder(w).Encode(resp)
		return
	}
}
