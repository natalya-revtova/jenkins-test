NAME	:= huffman

BIN_DIR		:= ./bin

.PHONY: all

install:
	go get -u "github.com/natalya-revtova/jenkins-test"

build:
	go build -o ${BIN_DIR}/${NAME}.exe ${NAME}.go

lint:
	golangci-lint run

test:
	go test -cover -coverprofile=coverage.txt

coverage:
	go tool cover -html=coverage.txt -o coverage.html

clean:
	go clean
	rm -f ${BIN_DIR}/${NAME}.exe