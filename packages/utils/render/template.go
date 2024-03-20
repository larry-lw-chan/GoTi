package render

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// Struct used to pass template path and layout location
type Location struct {
	TmplPath string
	Layout   map[string][]string
}

// Layout and Template Location
var layout map[string][]string
var tmplPath string

// Template Cache for storing parsed templates
var tmplCache = map[string]*template.Template{}

// Function used to declare the template path and layout location
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

/****************************************************
* render Template
****************************************************/
func Template(w http.ResponseWriter, data map[string]interface{}, tmplName string) {
	// get layout from filename Eg. home.base.tmpl -> base
	arr := strings.Split(tmplName, "/")
	filename := arr[len(arr)-1]
	layoutName := strings.Split(filename, ".")[1]

	// get template from cache
	tmpl, ok := tmplCache[tmplName]

	// if not available, parse and cache template
	if !ok {
		tmplFiles := getTmplFiles(tmplName, layoutName)
		tmpl = template.Must(template.ParseFiles(tmplFiles...))
		tmplCache[tmplName] = tmpl
	}

	// Execute template
	err := tmpl.ExecuteTemplate(w, layoutName, data)
	if err != nil {
		log.Println(err)
	}
}

func getTmplFiles(tmplName string, key string) []string {
	tmpl := tmplPath + tmplName
	tmplFiles := append(layout[key], tmpl)
	return tmplFiles
}

/****************************************************
* render Partial
****************************************************/
func Partial(w http.ResponseWriter, data map[string]interface{}, tmplName string) {
	// get template from cache
	tmpl, ok := tmplCache[tmplName]

	// if not available, parse and cache partial
	if !ok {
		tmplFile := tmplPath + tmplName
		tmpl = template.Must(template.ParseFiles(tmplFile))
		tmplCache[tmplName] = tmpl
	}
	// Execute partial
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
