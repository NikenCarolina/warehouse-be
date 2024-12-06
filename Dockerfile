# backend/Dockerfile
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/app/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/server /app/server

COPY .env /

EXPOSE 8080 

CMD ["/app/server"]
