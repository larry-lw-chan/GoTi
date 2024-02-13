package utils

import (
	"log"
	"net/http"
	"text/template"
)

var tmplPath = "./templates"
var tmplLayout = "./templates/layouts/base.layout.tmpl"

func Render(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl = tmplPath + tmpl
	t := template.Must(template.ParseFiles(tmpl, tmplLayout))
	err := t.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
