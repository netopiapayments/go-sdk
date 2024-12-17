.PHONY: lint test check

lint:
	golangci-lint run ./...

test: 
	go test -v ./...

check: lint test
	@echo "All checks passed."