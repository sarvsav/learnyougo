#!/usr/bin/env bash
if [[ $(uname) == "Linux" ]]; then
  docker run --rm --hostname learnyougo --name learnyougo -v "$PWD":/usr/src/learnyougo -w /usr/src/learnyougo golang:1.19 /bin/bash -c "go mod tidy; go build -v"
else
  export MSYS_NO_PATHCONV=1
  docker run --hostname learnyougo --name learnyougo --rm -v "$PWD":/usr/src/learnyougo -w /usr/src/learnyougo golang:1.19 /bin/bash -c "export GOOS=windows; export GOARCH=amd64; go mod tidy; go build -v"
fi
