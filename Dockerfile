FROM golang:1.11.5 AS build

WORKDIR /build
COPY . /build
RUN export GO111MODULE=on
RUN export GOOS=linux
RUN go build -o pgl .

FROM alpine:latest AS runtime

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /app
COPY --from=build /app/pgl ./

EXPOSE 6113:6113

CMD ["./pgl"]
