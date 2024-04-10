PHONY:
SILENT:

generate:
	go get -u github.com/99designs/gqlgen && go run github.com/99designs/gqlgen generate

build:
	go build -o ./.bin/main cmd/main/main.go

run: build 
	./.bin/main
