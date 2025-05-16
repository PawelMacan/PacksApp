FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/app .
COPY --from=builder /app/packs.json .

EXPOSE 8080

CMD ["./app"]
