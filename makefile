BUILD_DATE=$(shell date +%F)

build:
	go build -o angel ./src/client/angel/bin/*.go
