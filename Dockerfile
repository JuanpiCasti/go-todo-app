FROM golang:1.23.4-alpine3.21 AS builder

WORKDIR /app

ADD . .

RUN go build -o ./bin/serve ./cmd/serve.go

FROM alpine:3.21.2

WORKDIR /app

EXPOSE 8080

COPY --from=builder /app/bin/serve /app/serve

CMD ["./serve"]