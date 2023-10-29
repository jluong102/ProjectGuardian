BUILD_DATE=$(shell date +%F)
ANGEL_VERSION="v0.0.0"
LD_FLAGS=-X 'main.BUILD_DATE=$(BUILD_DATE)'

build:
	go build -ldflags "$(LD_FLAGS) -X 'main.VERSION=$(ANGEL_VERSION)'" -o angel ./src/client/angel/bin/*.go
