.PHONY: test
test:
	go test ./*

.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yml --timeout=15s ./...

.PHONY: install-migrate
migrate:
	go install github.com/rubenv/sql-migrate/...@latest

.PHONY: migrate
migrate:
	sql-migrate up -env="local"


