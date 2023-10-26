.PHONY: test

test:
	go test -v github.com/mebble/ltag/test

build:
	go build -o ltag ./cmd/ltag/main.go

bench: build
	sh ./test/benchmarks/run.sh
