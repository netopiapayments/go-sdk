.PHONY: lint check

lint:
	golangci-lint run ./...

check: lint
	@echo "All checks passed."