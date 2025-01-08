#!/usr/bin/env sh

if [ -f cmd/main/main.go ]; then
    echo "Build WASM application ...";
    GOOS=js GOARCH=wasm go build \
        -o application/hello.wasm \
        -trimpath cmd/main/*.go
else
    echo "wasm source not found!"
    exit 1
fi