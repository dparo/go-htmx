.DEFAULT_GOAL := all
.PHONY: all release clean tailwind run run-detached

all:
	go build


install-deps:
	go install github.com/cosmtrek/air@latest

tailwind:
	tailwindcss -i ./templates/input.css -o ./static/css/style.css

run: all
	./go-htmx

run-detached: all
	killall go-htmx || true
	./go-htmx &


serve: serve-watchman

serve-entr:
	find ./ static templates -not -path "./.git/*" -type f | entr -r make run
	# air

serve-watchman:
	watchman-make -p 'Makefile' '**/*.go' 'static/**' 'templates/**' '**.js' -t run-detached
