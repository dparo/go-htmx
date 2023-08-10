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


watch: watch-watchman

watch-air:
	air

watch-entr:
	find ./ static templates -not -path "./.git/*" -type f | entr -r make serve

watch-watchman:
	watchman-make -p 'Makefile' '**/*.go' 'web/**/*' '*.js' '*.json' -t serve-detached
