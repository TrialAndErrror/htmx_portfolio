.PHONY: build clean

# Default target
all: build

# Build the site
build:
	go run build.go

# Clean generated files
clean:
	rm -f index.html about.html projects.html wade.html

# Watch for changes and rebuild (requires fswatch or inotify-tools)
watch:
	@echo "Watching for changes... (Ctrl+C to stop)"
	@if command -v fswatch >/dev/null 2>&1; then \
		fswatch -o templates/ components/ build.py | xargs -n1 -I{} make build; \
	elif command -v inotifywait >/dev/null 2>&1; then \
		while inotifywait -r -e modify templates/ components/ build.py; do \
			make build; \
		done; \
	else \
		echo "Install fswatch or inotify-tools for file watching"; \
	fi

# Serve the site locally (requires Python's http.server)
serve: build
	python -m http.server 8000

# Build and serve
dev: build serve 