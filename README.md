# Trial and Errror Portfolio Website

A modern, component-based portfolio website built with Go templates and HTMX. This project features a clean build system that compiles templates into a production-ready `dist/` directory.

## Table of Contents

- [System Requirements](#system-requirements)
- [Installation](#installation)
- [Build Process](#build-process)
- [Templating System](#templating-system)
- [Project Structure](#project-structure)
- [Development Workflow](#development-workflow)
- [Deployment](#deployment)
- [Contributing](#contributing)

## System Requirements

### Prerequisites

- **Go** (version 1.21 or higher)
- **Git** (for version control)


## Build Process

### Overview

The build system compiles Go templates into static HTML files and organizes all assets into a production-ready `dist/` directory.

### Build Steps

1. **Directory Creation**: Creates `dist/` directory if it doesn't exist
2. **Template Discovery**: Automatically discovers all page templates in `templates/pages/`
3. **Component Inclusion**: Includes all component templates from `templates/pages/{pageName}/`
4. **Template Compilation**: Compiles all templates into HTML files in `dist/`
5. **Asset Copying**: Copies all static assets from `static/` to `dist/static/`

### Build Commands

```bash
# Build only (creates dist/ directory)
go run main.go 

# Build and start development server
go run main.go -serve

# Build and start server on custom port
go run main.go -serve -port=8080
```

### Build Output

After building, the `dist/` directory contains:

```
dist/
├── index.html          # Home page
├── about.html          # About page
├── projects.html       # Projects page
├── wade.html           # Wade's page
└── static/             # All static assets
    ├── css/            # Stylesheets
    ├── js/             # JavaScript files
    ├── images/         # Images and icons
    └── favicon.ico     # Favicon files
```

## Templating System

### Template Structure

The project uses Go's `html/template` package with a component-based architecture:

```
templates/
├── base.html                    # Base template with common layout
└── pages/                       # Page-specific templates
    ├── index.html              # Home page content
    ├── about.html              # About page content
    ├── projects.html           # Projects page content
    ├── wade.html               # Wade's page content
    └── projects/               # Component templates for projects page
        ├── cmg-inventory.html
        ├── pals-haven.html
        ├── flask-microservices.html
        ├── money-manager.html
        ├── cat-face-detector.html
        ├── tne-assistant.html
        └── fetcher.html
```

### Template Features

- **Base Template**: Common HTML structure, navigation, and styling
- **Page Templates**: Define the `content` block for each page
- **Component Templates**: Reusable sections that can be included in pages
- **Automatic Discovery**: Build system automatically finds and includes components

### Template Syntax

```html
<!-- Base template -->
{{define "base.html"}}
<!doctype html>
<html>
  <head>
    <!-- Common head content -->
  </head>
  <body>
    <!-- Navigation -->
    {{template "content" .}}
    <!-- Footer -->
  </body>
</html>
{{end}}

<!-- Page template -->
{{define "content"}}
<main>
  <h1>Page Title</h1>
  {{template "component-name" .}}
</main>
{{end}}

<!-- Component template -->
{{define "component-name"}}
<div class="component">
  <!-- Component content -->
</div>
{{end}}
```

## Project Structure

```
htmx_portfolio/
├── main.go                     # Build system and server
├── go.mod                      # Go module definition
├── .gitignore                  # Git ignore rules
├── README.md                   # This file
├── templates/                  # Source templates
│   ├── base.html              # Base template
│   └── pages/                 # Page templates
│       ├── index.html
│       ├── about.html
│       ├── projects.html
│       ├── wade.html
│       └── projects/          # Component templates
│           ├── cmg-inventory.html
│           ├── pals-haven.html
│           └── ...
├── static/                     # Source static assets
│   ├── css/                   # Stylesheets
│   ├── js/                    # JavaScript files
│   ├── images/                # Images and icons
│   └── favicon.ico            # Favicon files
├── pages/                      # Legacy pages (deprecated)
└── dist/                       # Build output (generated)
    ├── *.html                 # Compiled HTML files
    └── static/                # Copied static assets
```

## Deployment

### Local Development

```bash
# Build and start development server
go run main.go -serve

# Access the site at http://localhost:8000
```

### Production Deployment

1. Build the project: `go run main.go -serve=false`
2. Upload the contents of the `dist/` directory to your web server
3. Configure your web server to serve static files from the uploaded directory

### Web Server Configuration

#### Nginx Example

```nginx
server {
    listen 80;
    server_name your-domain.com;
    root /path/to/dist;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }

    location /static/ {
        expires 1y;
        add_header Cache-Control "public, immutable";
    }
}
```

#### Apache Example

```apache
<VirtualHost *:80>
    ServerName your-domain.com
    DocumentRoot /path/to/dist
    
    <Directory /path/to/dist>
        AllowOverride All
        Require all granted
    </Directory>
</VirtualHost>
```

## Contact
mailto:coding@trialanderrror.com
