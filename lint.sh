#!/bin/bash

set -ex

install(){
    echo "install lint tool"
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.52.2
    golangci-lint --version
#    integrations to your computer
#    https://golangci-lint.run/usage/integrations/
}

run(){
    golangci-lint run ./...
}

if [ "$1" = "install" ]
then
    install
elif [ "$1" = "run" ]
then
   run
else
    echo "command not found $1"
fi