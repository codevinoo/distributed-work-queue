.PHONY: build demo test clean

build:
	go build -o bin/broker ./cmd/broker
	go build -o bin/worker ./cmd/worker

demo:
	@./scripts/demo.sh

test:
	go test ./...

clean:
	rm -rf bin/
