.PHONY: test
.DEFAULT_GOAL := test

lint:
	golangci-lint run ./...

test:
	go test -v -cover -race -timeout 60s ./...

bench:
	go test -bench=. ./custommap_atomics/...
