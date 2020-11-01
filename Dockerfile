FROM golang:1.14

ADD . /go/src/github.com/harshshekhar15/monitoring

RUN go install /go/src/github.com/harshshekhar15/monitoring/cmd/server

ENTRYPOINT /go/bin/server

EXPOSE 8080