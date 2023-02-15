#!/bin/bash

cd "${GOPATH}/src/github.com/liangyt123/go-todo"

find . -name "*.go" -not -path "./vendor/*" | xargs gofmt -w

git diff --exit-code
