FROM golang:1.18-buster as builder

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_HOME /go/src/entify

WORKDIR "$APP_HOME"
ADD . "$APP_HOME"

RUN go mod download
RUN go mod verify
RUN go get ./...
RUN go mod vendor
RUN go build -o entify

EXPOSE 8010
CMD ["./entify"]