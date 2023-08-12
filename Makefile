.DEFAULT_GOAL := all
.PHONY: all release clean tailwind run run-detached

all: tailwind
	go build

install-deps:
	go install github.com/cosmtrek/air@latest

tailwind:
	tailwindcss -i ./web/templates/input.css -o ./web/static/css/style.css

serve: all
	./go-htmx

serve-detached: all
	killall go-htmx || true
	./go-htmx &


watch: watch-entr

watch-air:
	air

watch-entr:
	while true; do find . -type f \
			-name "Makefile" -o -name "*.go" -o -name "*.js" -o -name "*.css" -o -name "*.gohtml" -o -name "*.html" \
			! -path "./.git/*" ! -path "./web/static/css/style.css" | entr -r -d make serve; sleep 0.5; done


watch-watchman:
	watchman-make -p 'Makefile' '**/*.go' 'web/**/*' '*.js' '*.json' -t serve-detached
