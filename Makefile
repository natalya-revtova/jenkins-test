NAME	:= huffman

BIN_DIR		:= ./bin
OUT_DIR		:= ./reports

.PHONY: all

install:
	go get -u "github.com/natalya-revtova/jenkins-test"

build:
	go build -o ${BIN_DIR}/${NAME}.exe ${NAME}.go

lint:
	golangci-lint run

test:
	mkdir ${OUT_DIR}
	go test -coverprofile=${OUT_DIR}/coverage.txt -covermode=atomic

coverage:
	go tool cover -html=${OUT_DIR}/coverage.txt -o ${OUT_DIR}/coverage.html

clean:
	go clean
	rm -f ${BIN_DIR}/${NAME}.exe