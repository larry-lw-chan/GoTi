package render

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

const (
	tmplPath   = "./templates/"
	baseLayout = tmplPath + "layout/base.layout.tmpl"
	navPartial = tmplPath + "partial/nav.partial.tmpl"
)

// Todo: function will determine correct template subdirectory by parsing the file name
// Example: base.layout.tmpl equal 'layout' folder and 'base' file name
func getTmplFiles(tmplFile string) []string {
	tmpl := tmplPath + strings.Split(tmplFile, ".")[1] + "/" + tmplFile
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
