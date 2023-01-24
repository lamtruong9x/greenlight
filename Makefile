start:
	air -c .air.toml
lint:
	golangci-lint run ./...