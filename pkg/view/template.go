package view

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tpIndex = parseTemplate("template/root.tmpl", "template/index.tmpl")
)

func parseTemplate(files ...string) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{})
	_, err := t.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	t = t.Lookup("root")
	return t
}

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.Execute(w, data)
	if err != nil {
		log.Println(err)
		return
	}
}

// Index renders index view
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}
