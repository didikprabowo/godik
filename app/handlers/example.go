package handlers

import (
	templateku "gitlab.com/didik/godik/app/templates"

	// "encoding/json"
	"net/http"
)

type Bio struct {
	Name string
}

func Index(w http.ResponseWriter, r *http.Request) {
	data := []Bio{Bio{Name: "didik"}}
	templateku.AppTemplates(w, "hai.html", "Index", data)
}
