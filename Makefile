BIN := "./bin/collector"
BIN_CLIENT := "./bin/client"

build_gather:
	go build -v -o $(BIN) ./cmd/collector

run_gather:
	$(BIN) -config ./configs/config.yaml

build_client:
	go build -v -o $(BIN_CLIENT) ./cmd/client

run_client:
	$(BIN_CLIENT)

test:
	go test -race -count 100 ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.45.0

lint: install-lint-deps
	golangci-lint run ./...

generate:
	go generate ./...