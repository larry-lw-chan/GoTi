package render

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

// Structs
type Location struct {
	TmplPath string
	Layout   map[string][]string
}

// Function Variables
var layout map[string][]string
var tmplPath string

// Template Cache for storing parsed templates
var tmplCache = map[string]*template.Template{}

func Layouts(l Location) {
	tmplPath = l.TmplPath
	layout = l.Layout

	// Append Template Path to Layout
	for _, path := range layout {
		for i, file := range path {
			path[i] = tmplPath + file
		}
	}
}

// Templates
func Template(w http.ResponseWriter, tmplName string, data any) {
	// get layout from filename Eg. home.base.tmpl -> base
	filename := strings.Split(tmplName, "/")[1]
	layoutName := strings.Split(filename, ".")[1]

	// get template from cache or parse and cache it
	tmpl := getCacheTemplate(tmplName, layoutName)

	err := tmpl.ExecuteTemplate(w, layoutName, data)
	if err != nil {
		log.Println(err)
	}
}

// Get template from cache or parse and cache it
func getCacheTemplate(tmplName string, layoutName string) *template.Template {
	tmpl, ok := tmplCache[tmplName]

	if !ok {
		tmplFiles := getTmplFiles(tmplName, layoutName)
		tmpl = template.Must(template.ParseFiles(tmplFiles...))
		tmplCache[tmplName] = tmpl
	}

	return tmpl
}

func getTmplFiles(tmplFile string, key string) []string {
	tmpl := tmplPath + tmplFile
	tmplFiles := append([]string{tmpl}, layout[key]...)
	return tmplFiles
}
