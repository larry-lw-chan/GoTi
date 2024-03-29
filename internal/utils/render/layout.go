package render

import (
	"os"
	"strings"
)

// Function used to declare the template path and Layout location
var Layout map[string][]string = make(map[string][]string)
var TmplPath string
var LayoutFolder = "layout"

type Location struct {
	TmplPath     string
	LayoutFolder string
}

// Layout and partials automatically sorted into Layout Map
// by following the naming convention
func New(l Location) {
	TmplPath = l.TmplPath
	LayoutFolder = l.LayoutFolder

	// Get list of layout file names
	layoutDir := TmplPath + "/" + LayoutFolder + "/"
	files, _ := os.ReadDir(layoutDir)

	// Add the main layout files to the Layout Map
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "__") { // Skip Partial Files
			continue
		}

		// Add Layout Files to Layout Map
		key := strings.Split(file.Name(), ".")[0]
		value := layoutDir + file.Name()

		// Append Layout File to Layout Map
		Layout[key] = append(Layout[key], value)
	}

	// Add Partial Files to Layout Map
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), "__") { // Skip Main Files
			continue
		}

		// If the partial has a key, add it to the corresponding Layout[key]
		if len(strings.Split(file.Name(), ".")) == 3 {
			key := strings.Split(file.Name(), ".")[1]
			value := layoutDir + file.Name()
			Layout[key] = append(Layout[key], value)
		} else {
			// If the partial has no key, add it to all Layout since it's global
			value := layoutDir + file.Name()
			for key := range Layout {
				Layout[key] = append(Layout[key], value)
			}
		}
	}
}
