package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

type Template struct {
	tpl *template.Template
}

func NewTemplate(name string) Template {
	tpl, err := Parse(path.Join("templates", name))
	if err != nil {
		log.Fatalln(err)
	}

	return tpl 
}

func Parse(name string) (Template, error) {
	tpl, err := template.ParseFS(FS, name)
	if err != nil {
		return Template{}, fmt.Errorf("Error while parsing template %v", err)
	}

	return Template{tpl: tpl}, nil
}

func (t *Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := t.tpl.Execute(w, data)
	if err != nil {
		log.Printf("Error While Executing Template %v", err)
		http.Error(w, "Error While Executing Template", http.StatusInternalServerError)
		return
	}
}
