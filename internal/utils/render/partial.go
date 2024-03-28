package render

import (
	"html/template"
	"log"
	"net/http"
)

/****************************************************
* render Partial
****************************************************/
func Partial(w http.ResponseWriter, data map[string]interface{}, files ...string) {
	// get main template
	tmplName := files[0]

	// get template from cache
	tmpl, ok := tmplCache[tmplName]

	// if not available, parse and cache partial
	if !ok {
		// Compile file list
		var tmplFiles []string
		for _, file := range files {
			tmplFiles = append(tmplFiles, TmplPath+file)
		}

		tmpl = template.Must(template.ParseFiles(tmplFiles...))
		tmplCache[tmplName] = tmpl
	}

	// Execute partial
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
