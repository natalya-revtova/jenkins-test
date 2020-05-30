NAME	:= huffman

TEST_DIR	:= ./results
BIN_DIR		:= ./bin

.PHONY: all

install:
	go get -u "github.com/natalya-revtova/jenkins-test"

build:
	go build -o ${BIN_DIR}/${NAME}.exe ${NAME}.go

lint:
	golangci-lint run

test:
	go test -cover -coverprofile=${TEST_DIR}/coverage.out

coverage:
	go tool cover -html=${TEST_DIR}/coverage.out -o ${TEST_DIR}/coverage.html

clean:
	go clean
	rm -f ${BIN_DIR}/${NAME}.exe