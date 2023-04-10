build: 
	@go build -o ./bin/ai
run: build
	@./bin/ai $(P)
