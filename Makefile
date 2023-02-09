start:
	air -c .air.toml
lint:
	golangci-lint run ./...
run:
	go run ./cmd/api