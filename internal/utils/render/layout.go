package render

import (
	"os"
	"strings"
)

// Default path to the templates
var TmplPath string = "templates/default"
var LayoutFolder string = "layout"

// Layout Map
var Layout map[string][]string = make(map[string][]string)

// Location struct to define the path to the templates
type Options struct {
	TmplPath     string
	LayoutFolder string
}

// Layout/partials automatically sorted into Layout Map based on naming convention
func New(o Options) {
	// Set the TmplPath and LayoutFolder
	if o.TmplPath != "" {
		TmplPath = o.TmplPath
	}
	if o.LayoutFolder != "" {
		LayoutFolder = o.LayoutFolder
	}

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
