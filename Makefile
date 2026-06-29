.PHONY: test-all test-python test-node test-go

test-all: test-python test-node test-go

test-python:
	@echo "Running Python tests..."
	python3 -m unittest discover -s tests

test-node:
	@echo "Running Node.js tests..."
	npm test

test-go:
	@echo "Running Go tests..."
	go test -v ./...
