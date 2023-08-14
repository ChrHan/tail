
ALL: lint run

run:
	go run main.go -n 20 test.txt

lint:
	golangci-lint run

build:
	go build
