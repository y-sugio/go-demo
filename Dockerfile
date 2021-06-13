FROM golang:latest

ARG GO111MODULE=auto
ARG GOPATH=/go

ENV GO111MODULE=${GO111MODULE}
ENV GOPATH=${GOPATH}

COPY src ./src

CMD ["go", "run", "./src/server/server.go"]