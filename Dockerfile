FROM golang:alpine AS builder

WORKDIR $GOPATH/src/myapp 

COPY main.go .

RUN apk update && apk upgrade

RUN apk add git

RUN go get ./...

RUN go build -o /go/app

FROM alpine:latest

COPY --from=builder /go/app /go/app

CMD ["/go/app"] 

EXPOSE 8082
