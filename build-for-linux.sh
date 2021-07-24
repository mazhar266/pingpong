#!/bin/bash
docker build . -t pingpong:latest
docker run -it --rm -e CGO_ENABLED=0 --volume $PWD:/go/src/pingpong pingpong:latest go build -a -v
