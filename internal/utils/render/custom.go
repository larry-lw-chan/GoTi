package render

import "html/template"

/****************************************************
* Custom Template Functions
****************************************************/
var funcMap = template.FuncMap{
	"loadPartial": loadPartial,
}

func loadPartial(path string) string {
	return path
}
