.PHONY: test
test:
	go test ./*

.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yml --timeout=15s ./...

.PHONY: install-migrate
install-migrate:
	go install github.com/rubenv/sql-migrate/...@latest

.PHONY: migrate
migrate:
	sql-migrate up -env="local"

.PHONY: down-all
down-all:
	sql-migrate down -limit 100 -env="local"

