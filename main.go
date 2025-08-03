package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// PageConfig holds the configuration for each page
type PageConfig struct {
	Title       string
	Description string
	CSS         string
	ActiveNav   string
	PageName    string
}

// Pages configuration
var pages = map[string]PageConfig{
	"index": {
		Title:       "Trial and Errror - Home",
		Description: "Trial and Errror - Web Development, Technical Consulting, and Video Content Development. Learn about our services and projects.",
		CSS:         "index",
		ActiveNav:   "",
		PageName:    "index",
	},
	"about": {
		Title:       "Trial and Errror - About",
		Description: "About Trial and Errror - Learn about our web development services, technical consulting, and educational video content. Founded on the principle that everyone makes mistakes.",
		CSS:         "about",
		ActiveNav:   "about",
		PageName:    "about",
	},
	"projects": {
		Title:       "Trial and Errror - Projects",
		Description: "Projects by Trial and Errror - View our professional web applications, personal projects, and portfolio of work in web development and technology.",
		CSS:         "projects",
		ActiveNav:   "projects",
		PageName:    "projects",
	},
	"wade": {
		Title:       "Trial and Errror - Meet Wade",
		Description: "Meet Wade Green - Software Engineer, Attorney at Law, and Technology Educator. View Wade's professional experience, skills, and background.",
		CSS:         "wade",
		ActiveNav:   "wade",
		PageName:    "wade",
	},
}

func main() {
	// Parse command line flags
	serve := flag.Bool("serve", false, "Start HTTP server after building")
	port := flag.String("port", "8000", "Port for HTTP server")
	flag.Parse()

	if flag.NFlag() == 0 {
		// No flags detected, show help text
		prompts := []string{
			"Trial and Errror - Go Template Builder",
			"Used to build the site from templates.",
			"Flags:",
			"-serve: Start HTTP server after building",
			"-port: Provide port for HTTP server (default: 8000)",
		}

		for _, prompt := range prompts {
			fmt.Println(prompt)
		}
		return
	}

	// Always build before serving
	buildSite()
	// Start server if requested
	if *serve {
		startServer(*port)
	}
}

func buildSite() {
	// Discover pages from the pages folder
	discoveredPages := discoverPages()

	// Parse base template
	baseTemplate, err := template.ParseFiles("templates/base.html")
	if err != nil {
		log.Fatal("Error parsing base template:", err)
	}

	// Process each discovered page
	for pageName, config := range discoveredPages {
		// Parse the specific page template
		pageTemplatePath := filepath.Join("pages", pageName+".html")

		// Clone the base template and add the page template
		combinedTemplate, err := baseTemplate.Clone()
		if err != nil {
			log.Printf("Error cloning base template for %s: %v", pageName, err)
			continue
		}

		// Add the page template to the combined template
		combinedTemplate, err = combinedTemplate.ParseFiles(pageTemplatePath)
		if err != nil {
			log.Printf("Error adding page template for %s: %v", pageName, err)
			continue
		}

		// Create output file
		outputFile, err := os.Create(pageName + ".html")
		if err != nil {
			log.Printf("Error creating %s.html: %v", pageName, err)
			continue
		}
		defer outputFile.Close()

		// Execute the combined template
		err = combinedTemplate.ExecuteTemplate(outputFile, "base.html", config)
		if err != nil {
			log.Printf("Error executing template for %s: %v", pageName, err)
			continue
		}

		log.Printf("Built %s.html", pageName)
	}

	log.Println("Build complete!")
}

func discoverPages() map[string]PageConfig {
	discoveredPages := make(map[string]PageConfig)

	// Read the pages directory
	files, err := os.ReadDir("pages")
	if err != nil {
		log.Fatal("Error reading pages directory:", err)
	}

	// Process each HTML file in the pages directory
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".html") {
			pageName := strings.TrimSuffix(file.Name(), ".html")

			// Check if we have configuration for this page
			if config, exists := pages[pageName]; exists {
				config.PageName = pageName
				discoveredPages[pageName] = config
				log.Printf("Discovered page: %s", pageName)
			} else {
				log.Printf("Warning: No configuration found for page %s", pageName)
			}
		}
	}

	return discoveredPages
}

func startServer(port string) {
	serverAddress := fmt.Sprintf("http://localhost:%s", port)
	log.Printf("Starting server at %s", serverAddress)
	log.Printf("Press Ctrl+C to stop")

	// Create file server for current directory
	fs := http.FileServer(http.Dir("."))

	// Start server
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), fs); err != nil {
		log.Fatal("Server error:", err)
	}
}
