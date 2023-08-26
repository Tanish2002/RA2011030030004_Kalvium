package handlers

import (
	"kalvium/controllers"
	"net/http"
)

type Handler struct {
	Controller controllers.Controller
}

func (h *Handler) ExpressionHandler(w http.ResponseWriter, r *http.Request) {

}
