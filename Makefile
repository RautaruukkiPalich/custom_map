.PHONY: test
.DEFAULT_GOAL := test

lint:
	golangci-lint run ./...

test:
	go test -v -cover -race -timeout 10s ./...

testatom:
	go test -v -cover -race -timeout 10s ./custommap_atomics/...
	
vet:
	go vet .

bench:

