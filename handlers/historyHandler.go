package handlers

import (
	"html/template"
	"net/http"
	"time"
)

func (h *Handler) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	history, err := h.Controller.HistoryController()

	funcs := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"formatTime": func(timeStamp time.Time) string {
			format := "02/01/2006 03:04 PM"
			return timeStamp.Format(format)
		},
	}
	const tmplFile = "./views/history.html"
	tmpl := template.Must(template.New("history.html").Funcs(funcs).ParseFiles(tmplFile))

	// Set header to indicate the response is a html page
	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, history)
	if ok := h.errorResp(err, w); ok {
		return
	}
}

func add(x, y int) int {
	return x + y
}
