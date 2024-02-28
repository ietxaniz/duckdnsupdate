FROM golang:alpine as builder

LABEL maintainer="Iñigo Etxaniz <inigoetxaniz@gmail.com>"

WORKDIR /app

COPY main.go ./
COPY go.mod ./

RUN go build -ldflags="-w -s" -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]