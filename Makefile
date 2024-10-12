build:
	@go build -o bin/training

run: build
	@./bin/training

test:
	@go test ./... -v
