FROM golang:1.23.4-alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go build -o ./bin/serve ./cmd/serve.go

FROM alpine:3.21.2

WORKDIR /app

ARG SERVER_PORT
ENV SERVER_PORT=${SERVER_PORT}
EXPOSE ${SERVER_PORT}

COPY --from=builder /app/bin/serve /app/serve

CMD ["./serve"]