FROM golang:1.8
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
ADD . /go/src/github.com/shawncatz/go-destiny/service

ENV DATABASE_HOST=127.0.0.1
ENV DATABASE_NAME=destiny

ENTRYPOINT revel run github.com/shawncatz/go-destiny/service prod

EXPOSE 9000
