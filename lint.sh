#!/bin/sh
set -eux

GOLANGCILINT=$(pwd)/bin/golangci-lint

if [ ! -f "$GOLANGCILINT" ]; then
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.59.0
fi

"$GOLANGCILINT" run --timeout=5m --enable-all -D wrapcheck -D err113 -D gofumpt -D depguard
