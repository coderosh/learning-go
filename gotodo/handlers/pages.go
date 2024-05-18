package handlers

import (
	"net/http"

	"gotodo/views"
)

type PageHandler struct{}

func (h *PageHandler) GetHomePage() func(w http.ResponseWriter, r *http.Request) {
	tpl := views.NewTemplate("index.html")

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, map[string]string{
			"Name": "Your's ",
		})
	}
}
