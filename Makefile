
ALL: lint run

run:
	go run main.go test.txt

lint:
	golangci-lint run

build:
	go build
