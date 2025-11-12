#!/bin/bash
set -e

mkdir -p bin

# bash script to build for specific OS/ARCH
# usage: ./devbuild.sh [darwin-amd64|linux-amd64|windows-amd64]
if [ "$1" == "darwin-amd64" ]; then
    GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/coolsnippets-darwin-amd64
    exit 0
elif [ "$1" == "linux-amd64" ]; then
    GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/coolsnippets-linux-amd64
    exit 0
elif [ "$1" == "windows-amd64" ]; then
    GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/coolsnippets-windows-amd64.exe
    exit 0
fi

# If no argument is provided, build for the current OS
case "$(uname -s)" in
    Darwin)
        GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/coolsnippets-darwin-amd64
        ;;
    Linux)
        GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/coolsnippets-linux-amd64
        ;;
    MINGW*|MSYS*|CYGWIN*)
        GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/coolsnippets-windows-amd64.exe
        ;;
    *)
        echo "Unsupported OS"
        exit 1
        ;;
esac
