package render

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

/****************************************************
* Custom Template Functions
****************************************************/
var funcMap = template.FuncMap{}

/****************************************************
* Render Template and Partial
****************************************************/
var tmplCache = map[string]*template.Template{}

func Template(w http.ResponseWriter, data map[string]interface{}, files ...string) {
	// get main template
	tmplName := files[0]

	// get layout from filename Eg. home.base.tmpl -> base
	arr := strings.Split(tmplName, "/")
	filename := arr[len(arr)-1]
	layoutName := strings.Split(filename, ".")[1]

	// get template from cache
	tmpl, ok := tmplCache[tmplName]

	// if not available, parse and cache template
	if !ok {
		tmplFiles := getTmplFiles(layoutName, files)
		tmpl = template.Must(template.New(tmplName).Funcs(funcMap).ParseFiles(tmplFiles...))
		tmplCache[tmplName] = tmpl
	}

	// Execute template
	err := tmpl.ExecuteTemplate(w, layoutName, data)
	if err != nil {
		log.Println(err)
	}
}

func getTmplFiles(key string, files []string) []string {
	tmplFiles := Layout[key]
	for _, file := range files {
		tmplFiles = append(tmplFiles, TmplPath+file)
	}
	return tmplFiles
}

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
