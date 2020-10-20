FROM golang:alpine as builder

LABEL maintainer="Iñigo Etxaniz <inigoetxaniz@gmail.com>"

WORKDIR /app

COPY main.go ./

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .



FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
