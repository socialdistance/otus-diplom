build:
	go build -v -o ./bin/collector ./cmd/collector

run: build
	./bin/collector -config ./configs/config.yaml