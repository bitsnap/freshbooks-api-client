#!/usr/bin/env bash
set +eux

cd "$(dirname "$0")" || exit 1

golangci-lint run -c .golangci.yml --fix
go fmt ./...
gofmt -s -r 'interface{} -> any' -w .
goimports -w .
gofumpt -w .

which goimports > /dev/null || go install golang.org/x/tools/cmd/goimports@latest
which goimports > /dev/null || echo "Add \$GOPATH/bin to \$PATH" && exit 1
which gofumpt > /dev/null || go install mvdan.cc/gofumpt@latest
which ginkgo > /dev/null || go install github.com/onsi/ginkgo/v2/ginkgo@latest
