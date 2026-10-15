#!/bin/bash

set -euo pipefail

script_dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
lint_version=$(tr -d '\r\n' < "$script_dir/.golangci-lint-version")

case "${1:-}" in
    install)
        curl -fsSL "https://raw.githubusercontent.com/golangci/golangci-lint/$lint_version/install.sh" |
            sh -s -- -b "$(go env GOPATH)/bin" "$lint_version"
        golangci-lint --version
        ;;
    run)
        golangci-lint run --allow-serial-runners ./...
        ;;
    *)
        echo "Usage: $0 {install|run}" >&2
        exit 2
        ;;
esac
