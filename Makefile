export tag=latest

build:
	echo "building cleanup binary"
    GOOS=linux GOARCH=amd64 go build -o bin/amd64/cleanup
build-win:
	echo "building cleanup binary"
    GOOS=win GOARCH=amd64 go build -o bin/win/cleanup.exe

