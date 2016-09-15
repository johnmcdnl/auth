FROM golang:1.7
ADD . /go/src/github.com/johnmcdnl/auth
RUN go get -d
RUN go install github.com/johnmcdnl/auth
ENTRYPOINT /go/bin/auth
EXPOSE 8600