.PHONY: test
test:
	go test ./*

.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yml --timeout=15s ./...
