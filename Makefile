export GOOS=linux
export GOARCH=arm
APP=adt7410

.PHONY: build run clean
build: clean
	go build -o ${APP} main.go

run:
	go run -race main.go

clean:
	go clean
