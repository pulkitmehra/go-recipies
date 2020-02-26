build:
	go build -v -o bin/hello cmd/hello/main.go 

run: build
	bin/hello