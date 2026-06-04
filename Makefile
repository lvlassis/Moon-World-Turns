.PHONY: dev build

dev:
	go run main.go

build:
	mkdir -p build
	go build -o build/moon-world-turns main.go


