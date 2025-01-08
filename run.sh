#!/usr/bin/env sh

if [ -f server/main.go ]; then
    echo "Build Web Server ..." && go build \
    -o bin/webserver.exe \
    -trimpath server/main.go
else
    echo "server source not found!"
    exit 1
fi

if [ -f bin/webserver.exe ]; then
    export PORT="3000";\
        start http://localhost:3000;\
        exec bin/webserver.exe
else
    exit 1
fi