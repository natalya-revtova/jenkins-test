NAME	:= huffman

TEST_DIR	:= ./results
BIN_DIR		:= ./bin

build:
	go build -o ${BIN_DIR}/${NAME}.exe ${NAME}.go
.PHONY: build

lint:
	golangci-lint run
.PHONY: lint

test:
	go test -cover -coverprofile=${TEST_DIR}/coverage.out
.PHONY: test

coverage:
	go tool cover -html=${TEST_DIR}/coverage.out -o ${TEST_DIR}/coverage.html
.PHONY: coverage

clean:
	go clean
	rm -f ${BIN_DIR}/${NAME}.exe
.PHONY: clean