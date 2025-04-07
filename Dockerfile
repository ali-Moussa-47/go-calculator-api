
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .


RUN go build -o calculatorApi

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/calculatorApi .

EXPOSE 8080

CMD ["./calculatorApi"]
