#!/bin/bash
set -eu
BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
COMMIT=$(git rev-parse --short HEAD)
BUILT_AT=$(date +%FT%T%z)
BUILT_BY=$(git config user.name)

function go-build {
	set -x
	go build \
		-o build/emoji-cloner-$GOOS-$GOARCH${1:-} \
		-ldflags "-X main.commit=${COMMIT} -X main.builtAt=${BUILT_AT} -X main.builtBy=${BUILT_BY}" \
		main.go
	set +x
}

cd $BASEDIR
if [ -e build ]; then
	rm -rf build
fi
mkdir -p build

GOOS=darwin   GOARCH=amd64 go-build
GOOS=linux    GOARCH=amd64 go-build
GOOS=linux    GOARCH=arm   go-build
GOOS=windows  GOARCH=amd64 go-build .exe

