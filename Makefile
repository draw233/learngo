run: build
	@./bin/main
build: 
	@go build -o bin/learngo main.go