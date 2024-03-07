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

// Todo: Template caching
var tmplCache = map[string]*template.Template{}

// Shortcut for caching and rendering templates
func Template(w http.ResponseWriter, tmplName string, data any) {
	tmpl := getCacheTemplate(tmplName)

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

// Get template from cache or parse and cache it
func getCacheTemplate(tmplName string) *template.Template {
	tmpl, ok := tmplCache[tmplName]

	if !ok {
		tmplFiles := getTmplFiles(tmplName)
		tmpl = template.Must(template.ParseFiles(tmplFiles...))
		tmplCache[tmplName] = tmpl
	}

	return tmpl
}

func getTmplFiles(tmplFile string) []string {
	tmpl := tmplPath + tmplFile
	tmplFiles := []string{tmpl, baseLayout, navPartial}
	return tmplFiles
}
