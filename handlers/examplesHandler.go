package handlers

import (
	"html/template"
	"net/http"
)

func (h *Handler) ExamplesHandler(w http.ResponseWriter, r *http.Request) {
	const tmplFile = "./views/example.html"
	tmpl := template.Must(template.New("example.html").ParseFiles(tmplFile))
	err := tmpl.Execute(w, nil)
	if ok := h.errorResp(err, w); ok {
		return
	}
}
