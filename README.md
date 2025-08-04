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


## Installation

### Debian/Ubuntu

```bash
# Update package list
sudo apt update

# Install Go
sudo apt install golang-go

# Install Git
sudo apt install git

# Verify installation
go version
git --version
```

### Arch 

```bash
# Update package list
sudo pacman -Syu

# Install Go
sudo pacman -S go

# Install Git
sudo pacman -S git

# Verify installation
go version
git --version
```

### Clone and Setup

```bash
# Clone the repository
git clone <repository-url>
cd htmx_portfolio

# Initialize Go modules (if not already done)
go mod init htmx_portfolio

# Build the project
go run main.go 
```

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

## Development Workflow

### Adding New Pages

1. Create a new template file in `templates/pages/`
2. Define the `content` block with your page content
3. Add page configuration to `main.go` if needed
4. Build the project to generate the new page

### Adding New Components

1. Create a new component file in `templates/pages/{pageName}/`
2. Define a unique template name for your component
3. Include the component in your page template using `{{template "component-name" .}}`
4. Build the project to include the new component

### Modifying Existing Content

1. Edit the appropriate template file
2. Build the project to regenerate the HTML
3. Test your changes locally

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

## Contributing

### Development Setup

1. Fork the repository
2. Clone your fork locally
3. Create a feature branch
4. Make your changes
5. Test your changes locally
6. Submit a pull request

### Code Style

- Follow Go naming conventions
- Use meaningful template names
- Keep components focused and reusable
- Document any complex template logic

### Testing

- Always test your changes locally before submitting
- Verify that all pages build correctly
- Check that static assets are properly served
- Test responsive design on different screen sizes

## License

[Add your license information here]

## Contact

[Add your contact information here] 