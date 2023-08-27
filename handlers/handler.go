package handlers

import (
	"encoding/json"
	"kalvium/controllers"
	"net/http"
)

type Handler struct {
	Controller controllers.Controller
}

type errorResponse struct {
	Error string `json:"error"`
}

func (h *Handler) errorResp(err error, w http.ResponseWriter) bool {
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := errorResponse{
			Error: err.Error(),
		}
		err = json.NewEncoder(w).Encode(resp)
		return true
	}
	return false
}
