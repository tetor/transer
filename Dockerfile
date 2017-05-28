FROM golang:1.8

ENV GOPATH /home/go
ENV PATH="$PATH:$GOPATH/bin"
ENV WD $GOPATH/src/github.com/tetor/transer

WORKDIR $WD

RUN apt-get update

RUN go get github.com/spf13/cobra/cobra

CMD ["go", "run", "main.go"]
