.PHONY: build
build:
	go build -v ./cmd/clnkserver

.PHONY: serve
serve:
	yarn --cwd web serve

.DEFAULT_GOAL := build