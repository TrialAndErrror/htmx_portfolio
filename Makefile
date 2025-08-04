.PHONY: help build clean watch serve dev

# Default target - show help
.DEFAULT_GOAL := help

# Show this help message
help:
	@echo "🚀 Trial and Errror - Static Site Generator"
	@echo ""
	@echo "Available commands:"
	@echo ""
	@echo "  make build    Build the site from templates"
	@echo "  make clean    Remove generated HTML files"
	@echo "  make watch    Watch for changes and auto-rebuild"
	@echo "  make serve    Build and serve locally on port 8000"
	@echo "  make dev      Build and serve (development mode)"
	@echo ""
	@echo "Examples:"
	@echo "  make          Show this help message"
	@echo "  make build    Generate index.html, about.html, etc."
	@echo "  make dev      Build site and open http://localhost:8000"
	@echo ""
	@echo "File structure:"
	@echo "  templates/    Template files (base.html, index.html, etc.)"
	@echo "  build.go      Go build script"
	@echo "  *.html        Generated HTML files (after running make build)"

# Build the site
build:
	@echo "🔨 Building site..."
	go run main.go -build
	@echo "✅ Build complete!"

# Clean generated files
clean:
	@echo "🧹 Cleaning generated files..."
	rm -f index.html about.html projects.html wade.html
	@echo "✅ Clean complete!"

# Watch for changes and rebuild (requires fswatch or inotify-tools)
watch:
	@echo "👀 Watching for changes... (Ctrl+C to stop)"
	@if command -v fswatch >/dev/null 2>&1; then \
		fswatch -o templates/ main.go static/ pages/  | xargs -n1 -I{} make build; \
	elif command -v inotifywait >/dev/null 2>&1; then \
		while inotifywait -r -e modify templates/ main.go static/ pages/; do \
			make build; \
		done; \
	else \
		echo "❌ Install fswatch or inotify-tools for file watching"; \
		echo "   macOS: brew install fswatch"; \
		echo "   Linux: sudo apt-get install inotify-tools"; \
	fi

# Serve the site locally using Go's built-in server
serve: 
	@echo "🌐 Starting local server at http://localhost:8000"
	@echo "   Press Ctrl+C to stop"
	go run main.go -serve

# Build and serve (development mode)
dev: 
	@echo "🚀 Starting development server..."
	@echo "   Site: http://localhost:8000"
	@echo "   Press Ctrl+C to stop"
	go run main.go -serve