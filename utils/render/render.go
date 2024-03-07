package render

import (
	"log"
	"net/http"
	"text/template"
)

const (
	tmplPath   = "./templates/"
	baseLayout = tmplPath + "layout/base.layout.html"
	navPartial = tmplPath + "partial/nav.partial.html"
)

// Todo: function will determine correct template subdirectory by parsing the file name
// Example: base.layout.tmpl equal 'layout' folder and 'base' file name
func getTmplFiles(tmplFile string) []string {
	tmpl := tmplPath + tmplFile
	tmplFiles := []string{tmpl, baseLayout, navPartial}
	return tmplFiles
}

func Template(w http.ResponseWriter, tmplFile string, data any) {
	tmplFiles := getTmplFiles(tmplFile)
	tmpl := template.Must(template.ParseFiles(tmplFiles...))

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
