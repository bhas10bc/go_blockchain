build:
	go build -o ./bin/anto

run: build
	./bin/anto

test:
	go test -v ./...