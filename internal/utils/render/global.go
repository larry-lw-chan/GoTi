package render

import (
	"html/template"
)

// Struct used to pass template path and layout location
type Location struct {
	TmplPath string
	Layout   map[string][]string
}

// Layout and Template Location
var Layout map[string][]string
var TmplPath string

// Template Cache for storing parsed templates
var tmplCache = map[string]*template.Template{}

// Function used to declare the template path and Layout location
func Layouts(l Location) {
	TmplPath = l.TmplPath
	Layout = l.Layout

	// Append Template Path to Layout
	for _, path := range Layout {
		for i, file := range path {
			path[i] = TmplPath + file
		}
	}
}
