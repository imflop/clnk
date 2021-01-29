.PHONY: run
run:
	go run --work ./cmd/clnkserver

.PHONY: build
build:
	go build -v ./cmd/clnkserver

.PHONY: serve
serve:
	yarn --cwd web serve

.PHONY: vue-build
vue-build:
	yarn --cwd web build

.DEFAULT_GOAL := build