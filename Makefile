test:
	go test -v github.com/mebble/ltag/src

bench:
	sh ./benchmarks/run.sh

build:
	go build -o ltag
