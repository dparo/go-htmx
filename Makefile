.DEFAULT_GOAL := all
.PHONY: all release clean

all:
	echo Hello world

install-deps:
	go install github.com/cosmtrek/air@latest

serve:
	air
