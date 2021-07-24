#!/bin/bash
docker build . -t pingpong:latest
docker run -it --rm --volume $PWD:/go/src/pingpong pingpong:latest go build main.go
