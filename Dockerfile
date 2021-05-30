FROM golang:latest

RUN mkdir /go-demo
COPY src /go-demo/src

CMD ["go", "run", "/go-demo/src/main/main.go"]