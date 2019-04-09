FROM golang:1.11.5-alpine3.9

WORKDIR /build
COPY . /build
RUN export GO111MODULE=on
RUN export GOOS=linux
RUN go build -o pgl .

EXPOSE 6113:6113

CMD ["./pgl"]
