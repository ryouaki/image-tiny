# rustup target list
all: init

init:
	./build.sh aarch64-apple-darwin

build:
	./build.sh aarch64-apple-darwin
	go build -ldflags "-s -w" -o ./bin/imagetiny ./main.go 