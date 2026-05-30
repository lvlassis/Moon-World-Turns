.PHONY: dev build

dev:
	go run main.go

build:
	mkdir -p build
	go build -o build/main main.go


