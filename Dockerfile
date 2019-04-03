FROM golang:1.11.5

WORKDIR /build
COPY . /build
RUN export GO111MODULE=on
RUN export GOOS=linux
RUN go build -o pgl .

ENTRYPOINT ./pgl

EXPOSE 6113:6113
ENV PORT=3333
