FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /main .

COPY .env .

CMD ["./main"]

