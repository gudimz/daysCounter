BINARY_NAME=daysCounter

build:
	go build -o ${BINARY_NAME} main.go

clean:
	go clean