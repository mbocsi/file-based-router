build:
	@go build -o ./bin/fbr 

run: build
	@./bin/fbr

test:
	@go test ./... -v
