FROM golang:1.14
WORKDIR /go/src/app
RUN go get github.com/codegangsta/gin
COPY ./src .
EXPOSE 3000
CMD ["gin","run","server.go"]
