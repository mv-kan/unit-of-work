FROM golang:1.18-alpine

WORKDIR /dockerapp

COPY . . 

RUN go mod vendor

RUN go build -o /main

CMD ["/main"]