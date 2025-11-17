FROM golang:1.25.1-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
WORKDIR /app/cmd/service
RUN go build -o /app/bank main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bank .
COPY --from=builder /app/testdb.csv .
CMD ["./bank"]