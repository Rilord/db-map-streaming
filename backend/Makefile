BINARY_NAME=app.out

all: build test

build:
	go build -o ${BINARY_NAME} app.go

test:
	go test -v app.go


run:
	go build -o ${BINARY_NAME} app.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
