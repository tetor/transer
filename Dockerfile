FROM golang:1.8

ENV GOPATH /home/.go
ENV PATH="$PATH:$GOPATH/bin"

RUN apt-get update

RUN go get github.com/spf13/cobra/cobra
