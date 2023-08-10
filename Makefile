.DEFAULT_GOAL := all
.PHONY: all release clean

all: tailwind


install-deps:
	go install github.com/cosmtrek/air@latest

tailwind:
	tailwindcss -i ./templates/input.css -o ./static/css/style.css

run:
	go build
	./go-htmx

serve:
	find ./ static templates -not -path "./.git/*" -type f | entr -r make run
	# air
