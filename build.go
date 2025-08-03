package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

// PageConfig holds the configuration for each page
type PageConfig struct {
	Title       string
	Description string
	CSS         string
	ActiveNav   string
}

// Pages configuration
var pages = map[string]PageConfig{
	"index": {
		Title:       "Trial and Errror - Home",
		Description: "Trial and Errror - Web Development, Technical Consulting, and Video Content Development. Learn about our services and projects.",
		CSS:         "index",
		ActiveNav:   "",
	},
	"about": {
		Title:       "Trial and Errror - About",
		Description: "About Trial and Errror - Learn about our web development services, technical consulting, and educational video content. Founded on the principle that everyone makes mistakes.",
		CSS:         "about",
		ActiveNav:   "about",
	},
	"projects": {
		Title:       "Trial and Errror - Projects",
		Description: "Projects by Trial and Errror - View our professional web applications, personal projects, and portfolio of work in web development and technology.",
		CSS:         "projects",
		ActiveNav:   "projects",
	},
	"wade": {
		Title:       "Trial and Errror - Meet Wade",
		Description: "Meet Wade Green - Software Engineer, Attorney at Law, and Technology Educator. View Wade's professional experience, skills, and background.",
		CSS:         "wade",
		ActiveNav:   "wade",
	},
}

func main() {
	// Parse templates
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Error parsing templates:", err)
	}

	// Process each page
	for pageName, config := range pages {
		// Create output file
		outputFile, err := os.Create(pageName + ".html")
		if err != nil {
			log.Printf("Error creating %s.html: %v", pageName, err)
			continue
		}
		defer outputFile.Close()

		// Execute template
		err = tmpl.ExecuteTemplate(outputFile, pageName+".html", config)
		if err != nil {
			log.Printf("Error executing template for %s: %v", pageName, err)
			continue
		}

		log.Printf("Built %s.html", pageName)
	}

	log.Println("Build complete!")
} 