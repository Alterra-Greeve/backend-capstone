FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go mod verify

RUN go mod tidy

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app .

CMD ["./main"]
