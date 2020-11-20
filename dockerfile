FROM golang:latest
WORKDIR $GOPATH/go-web-hello-world
ADD . $GOPATH/go-web-hello-world
RUN go build .
EXPOSE 8082
ENTRYPOINT  ["./main"]
