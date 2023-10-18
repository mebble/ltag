# when we use `go run .` instead of building the binary, the "broken pipe" error appears
seq 100000000 | go run . | head
