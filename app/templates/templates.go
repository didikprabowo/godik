package template

import (
	"html/template"
	"net/http"
	"path"
)

func AppTemplates(w http.ResponseWriter, name string, varians string, data interface{}) {

	if name == "" {
		name = "default-a"
	}
	if varians == "" {
		name = "default-a"
	}
	var filepath = path.Join("app/templates/views", name)

	t, err := template.ParseFiles(filepath)
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(w, varians, data)
	if err != nil {
		panic(err)
	}

}
