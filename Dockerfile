FROM golang:1.11.5-alpine3.9

RUN apk update && apk add git
RUN apk add --no-cache ca-certificates cmake make g++ openssl-dev

WORKDIR /build
COPY . /build
RUN export GO111MODULE=on
RUN export GOOS=linux
RUN go build -o pgl .

EXPOSE 6113:6113


CMD ["./pgl"]
