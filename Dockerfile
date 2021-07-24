# grab from go official image
FROM golang:1.16-alpine

# make working dir
RUN mkdir src/pingpong
WORKDIR /go/src/pingpong

# copy the source
COPY . .
RUN go get .
RUN go build main.go

CMD ["main"]
