package render

// Function used to declare the template path and Layout location
var Layout map[string][]string
var TmplPath string

type Location struct {
	TmplPath string
	Layout   map[string][]string
}

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
